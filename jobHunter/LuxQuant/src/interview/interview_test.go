package interview

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"
)

func TestKlinesPub(t *testing.T) {
	t.Logf("begin...")
	file, err := os.Open("../dataSample.txt")
	if err != nil {
		t.Errorf("fail:%s", err)
		log.Fatal(err)
	}
	defer file.Close()

	testSub := KlinesPub{
		firstKLine: true,
	}

	symbolsSub := []string{"ETH-USDT", "BTC-USDT"}

	// sub 标的物
	testSub.Init(symbolsSub)

	// 启动协程，定时在每分钟开始的500ms时刻进行Pub
	// 使用实时数据源
	// go testSub.PublishTask()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// 接收到行情 或 交易信息，进行处理
		//t.Logf("%s", line)
		parts := strings.Split(line, " ")

		if parts[0] == "OnTrade" {
			var event Trade
			if err := json.Unmarshal([]byte(parts[1]), &event); err != nil {
				t.Logf("json.Unmarshal fail, OnTrade string:%s", parts[1])
			}
			testSub.OnTrade(event)
		}

		if parts[0] == "OnTick" {
			var event Tick
			if err := json.Unmarshal([]byte(parts[1]), &event); err != nil {
				t.Logf("json.Unmarshal fail, OnTick string:%s", parts[1])
			}
			testSub.OnTick(event)
		}
	}

	// 使用测试数据
	t.Log("-->Publish")
	testSub.Publish()
	testSub.Publish()

	if err := scanner.Err(); err != nil {
		t.Log(err)
	}
}
