package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 1. 创建 server
	s := server.NewMCPServer(
		"binance-price-server", // 名字
		"0.1.0",                // 版本
		server.WithToolCapabilities(true),
	)

	// 2. 定义一个 tool
	priceTool := mcp.NewTool("get_binance_price",
		mcp.WithDescription("查询币安某个交易对的最新价格"),
		mcp.WithString("symbol",
			mcp.Required(),
			mcp.Description("交易对，例如 BTCUSDT、ETHUSDT"),
		),
	)

	// 3. 注册 handler
	s.AddTool(priceTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		symbol, err := req.RequireString("symbol")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		resp, err := http.Get(fmt.Sprintf(
			"https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		defer resp.Body.Close()

		var data struct {
			Symbol string `json:"symbol"`
			Price  string `json:"price"`
		}
		json.NewDecoder(resp.Body).Decode(&data)

		return mcp.NewToolResultText(
			fmt.Sprintf("%s 当前价格：%s USDT", data.Symbol, data.Price),
		), nil
	})

	// 4. 启动（stdio 模式）
	if err := server.ServeStdio(s); err != nil {
		panic(err)
	}
}
