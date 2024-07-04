package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"future-go/go-zero_demo/micro_mall/gate/internal/config"
	"future-go/go-zero_demo/micro_mall/gate/internal/errorx"
	"future-go/go-zero_demo/micro_mall/gate/internal/handler"
	"future-go/go-zero_demo/micro_mall/gate/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/gate-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义异常 处理为 json 形式
	// 处理的传入err已经是interface了，需要将其转化为具体的结构体形式
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (httpStatus int, errorData any) {
		switch e := err.(type) {
		case *errorx.BizError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
