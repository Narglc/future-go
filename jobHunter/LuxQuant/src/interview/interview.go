package interview

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/pebbe/zmq4"
)

type Tick struct {
	Symbol       string  `json:"symbol"`
	BestBidPrice float64 `json:"bbp"`
	BestBidSize  float64 `json:"bbs"`
	BestAskPrice float64 `json:"bap"`
	BestAskSize  float64 `json:"bas"`
	Timestamp    int64   `json:"timestamp"`
}

type Trade struct {
	Symbol     string  `json:"symbol"`
	Price      float64 `json:"price"`
	Size       float64 `json:"size"`
	TradeCount int64   `json:"tradecount"`
	IsBuy      bool    `json:"m"`
	Timestamp  int64   `json:"timestamp"`
}

type KlineJson struct {
	Symbol    string  `json:"symbol"`
	OpenTime  int64   `json:"open_time"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Twap      float64 `json:"twap"`
	SellCount int64   `json:"sell_count"`
	SellSize  float64 `json:"sell_size"`
	BuyCount  int64   `json:"buy_count"`
	BuySize   float64 `json:"buy_size"`
}

type KlinesMessage struct {
	Klines map[string]KlineJson `json:"klines"` //key: symbol
}

type KlinesPub struct {
	firstKLine  bool          // 是否为程序启动后第一根k线
	symbols     []string      // sub的标的物列表
	pubMsg      KlinesMessage // 待pub的Kline
	curKLineMsg KlinesMessage // 当前正在统计的分钟线数据
	publisher   *zmq4.Socket  // zeromq 发布端socket
}

func initZeroMq() *zmq4.Socket {
	// 创建一个ZeroMQ上下文
	context, err := zmq4.NewContext()

	if err != nil {
		log.Fatalf("zmq4.NewContext fail, err:%v", err)
	}

	// 创建PUB套接字
	publisher, err := context.NewSocket(zmq4.PUB)
	if err != nil {
		log.Fatalf("context.NewSocket fail, err:%v", err)
	}

	// 绑定PUB套接字到指定地址
	publisher.Bind("tcp://127.0.0.1:5001")
	return publisher
}

func (s *KlinesPub) Init(symbols []string) {
	s.firstKLine = true
	s.symbols = symbols
	s.pubMsg.Klines = make(map[string]KlineJson)
	s.curKLineMsg.Klines = make(map[string]KlineJson)
	s.publisher = initZeroMq()
}

func (s *KlinesPub) OnTick(event Tick) {
	// 只处理sub的标的物
	isSub := false
	for _, each := range s.symbols {
		if each == event.Symbol {
			isSub = true
			break
		}
	}

	if !isSub {
		//fmt.Printf("未订阅该标的：%s", event.Symbol)
		return
	}

	currentCandlestick, ok := s.curKLineMsg.Klines[event.Symbol]
	if !ok {
		currentCandlestick = KlineJson{}
	}

	price := (event.BestBidPrice + event.BestAskPrice) / 2
	curTickTime := time.Unix(event.Timestamp/1e9, 0)
	fmt.Printf("OnTick:%s:-->Time:%v\n", event.Symbol, curTickTime)

	// 第一次行情数据
	if currentCandlestick.Open == 0 {
		currentCandlestick = KlineJson{
			Symbol:   event.Symbol,
			OpenTime: event.Timestamp,
			Open:     price,
			High:     price,
			Low:      price,
			Close:    price,
			Twap:     price,
			SellSize: event.BestAskSize,
			BuySize:  event.BestBidSize,
		}
		s.curKLineMsg.Klines[event.Symbol] = currentCandlestick

	} else if curTickTime.Minute() != time.Unix(currentCandlestick.OpenTime/1e9, 0).Minute() {
		//TODO: twap 如何计算？待给出公式

		// 第二分钟行情来了,将上一分钟的数据移到 s.pubMsg
		s.pubMsg.Klines[event.Symbol] = currentCandlestick

		newCandlestick := KlineJson{
			Symbol:   event.Symbol,
			OpenTime: event.Timestamp,
			Open:     price,
			High:     price,
			Low:      price,
			Close:    price,
			Twap:     price,
			SellSize: event.BestAskSize,
			BuySize:  event.BestBidSize,
		}
		currentCandlestick = newCandlestick
	} else {
		// 更新最高价，最低价，分钟收尾价
		currentCandlestick.Close = price

		if price > currentCandlestick.High {
			currentCandlestick.High = price
		}
		if price < currentCandlestick.Low {
			currentCandlestick.Low = price
		}

		currentCandlestick.SellSize += event.BestAskSize
		currentCandlestick.BuySize += event.BestBidSize
	}
}

func (s *KlinesPub) OnTrade(event Trade) {
	// 只处理sub的标的物
	isSub := false
	for _, each := range s.symbols {
		if each == event.Symbol {
			isSub = true
			break
		}
	}

	if !isSub {
		//fmt.Printf("未订阅该标的：%s", event.Symbol)
		return
	}

	currentCandlestick, ok := s.curKLineMsg.Klines[event.Symbol]
	if !ok {
		currentCandlestick = KlineJson{}
	}

	curTickTime := time.Unix(event.Timestamp/1e9, 0)
	fmt.Printf("OnTick:%s:-->Time:%v\n", event.Symbol, curTickTime)

	// 第一次行情数据
	if currentCandlestick.Open == 0 {
		currentCandlestick = KlineJson{
			Symbol:   event.Symbol,
			OpenTime: event.Timestamp,
			Open:     event.Price,
			High:     event.Price,
			Low:      event.Price,
			Close:    event.Price,
			Twap:     event.Price,
		}
		if event.IsBuy {
			currentCandlestick.BuySize = event.Size
			currentCandlestick.BuyCount = event.TradeCount
		} else {
			currentCandlestick.SellSize = event.Size
			currentCandlestick.SellCount = event.TradeCount
		}
		s.curKLineMsg.Klines[event.Symbol] = currentCandlestick

	} else if curTickTime.Minute() != time.Unix(currentCandlestick.OpenTime/1e9, 0).Minute() {
		// twap 计算
		if event.IsBuy {
			currentCandlestick.Twap = currentCandlestick.BuySize / float64(currentCandlestick.BuyCount)
		} else {
			currentCandlestick.Twap = currentCandlestick.SellSize / float64(currentCandlestick.SellCount)
		}

		// 第二分钟行情来了,将上一分钟的数据移到 s.pubMsg
		s.pubMsg.Klines[event.Symbol] = currentCandlestick

		newCandlestick := KlineJson{
			Symbol:   event.Symbol,
			OpenTime: event.Timestamp,
			Open:     currentCandlestick.Close, // k线的open为上一根线的close
			High:     event.Price,
			Low:      event.Price,
			Close:    event.Price,
			Twap:     event.Price,
		}
		if event.IsBuy {
			currentCandlestick.BuySize = event.Size
			currentCandlestick.BuyCount = event.TradeCount
		} else {
			currentCandlestick.SellSize = event.Size
			currentCandlestick.SellCount = event.TradeCount
		}

		currentCandlestick = newCandlestick
	} else {
		// 更新最高价，最低价，分钟收尾价
		currentCandlestick.Close = event.Price

		if event.Price > currentCandlestick.High {
			currentCandlestick.High = event.Price
		}
		if event.Price < currentCandlestick.Low {
			currentCandlestick.Low = event.Price
		}

		if event.IsBuy {
			currentCandlestick.BuySize += event.Size
			currentCandlestick.BuyCount += event.TradeCount
		} else {
			currentCandlestick.SellSize += event.Size
			currentCandlestick.SellCount += event.TradeCount
		}
	}
}

func (s *KlinesPub) Publish() {
	//测试出口
	s.pubMsg.Klines = s.curKLineMsg.Klines

	// 程序启动后第一根k线丢弃掉，因为不完整
	if s.firstKLine {
		s.firstKLine = false
		return
	}

	// 检查所有sub是否都已经汇总了分钟线
	if len(s.symbols) != len(s.pubMsg.Klines) {
		fmt.Printf("未收集完成,请检查数据源")
	}

	out, err := json.Marshal(s.pubMsg.Klines)
	if err != nil {
		fmt.Printf("pub err:%s", err)
	}
	fmt.Printf("KlinesPub Json:%s\n", out)

	// pub到zeromq
	if n, err := s.publisher.Send(string(out), 0); err != nil {
		fmt.Printf("pub.Send err:%v\n", err)
	} else {
		fmt.Printf("pub.Send %d bytes.\n", n)
	}

	// pub后需要将pubMsg数据清空
	s.pubMsg.Klines = make(map[string]KlineJson)
}

func (s *KlinesPub) PublishTask() {
	for {
		// 执行Pub任务，第一次会被丢弃
		s.Publish()

		// 获取当前时间
		now := time.Now()

		// 计算下一分钟的起始时间
		nextMinute := now.Add(time.Minute).Truncate(time.Minute)

		// 计算下一分钟的起始时间后的500毫秒
		targetTime := nextMinute.Add(500 * time.Millisecond)

		// 等待到目标时间
		sleepDuration := targetTime.Sub(now)
		time.Sleep(sleepDuration)
	}
}
