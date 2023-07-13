package interview

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pebbe/zmq4"
)

func zmqRecv() {
	subscriber, err := zmq4.NewSocket(zmq4.SUB)

	if err != nil {
		log.Fatalf("zmq4.NewContext fail, err:%v", err)
	}

	defer subscriber.Close()

	err = subscriber.Connect("tcp://127.0.0.1:5002") // 连接到Pub（发布）端的地址

	if err != nil {
		log.Fatalf("subscriber.Connect fail, err:%v", err)
	}

	err = subscriber.SetSubscribe("") // 订阅所有消息

	if err != nil {
		log.Fatalf("subscriber.SetSubscribe fail, err:%v", err)
	}

	for {
		msg, _ := subscriber.Recv(0)
		fmt.Println("zmq subscriber Received:", msg)
	}
}
func TestKlinesPub(t *testing.T) {
	file, err := os.Open("../dataSample.txt")
	if err != nil {
		t.Errorf("fail:%s", err)
		log.Fatal(err)
	}
	defer file.Close()

	testSub := KlinesPub{
		firstKLine: true,
	}

	symbolsSub := []string{"ETH-USDT", "APT-USDT"}

	// sub 标的物
	testSub.Init(symbolsSub)

	// 启动协程，定时在每分钟开始的500ms时刻进行Pub
	// 使用实时数据源
	go testSub.Publish()

	// 启动协程模拟zmq订阅端
	go zmqRecv()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// 接收到行情 或 交易信息，进行处理
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

	// 使用测试数据,dataSample数据量不足，主动打印
	// t.Log("-->Publish")
	// testSub.Publish()
	// testSub.Publish()

	if err := scanner.Err(); err != nil {
		t.Log(err)
	}

	time.Sleep(time.Second * 120)
}
