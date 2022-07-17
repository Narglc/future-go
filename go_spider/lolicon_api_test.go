package main

import (
	"fmt"
	sp "future-go/go_spider/fuliba_sort"
	"log"
	"testing"
	"time"
)

func TestLoliConApi(t *testing.T) {
	url := "https://api.lolicon.app/setu/v2"

	rsp := sp.GetResponse(url, nil, nil)

	log.Print((rsp.Body))
	/*
	   var result Result
	   _ = json.Unmarshal(bodyText, &result)
	   fmt.Println(i,result.Data)
	*/

}

func TestSimple(t *testing.T) {
	sl := []int{23, 56}
	for k, v := range sl {
		fmt.Print("--->", k, v)
	}
	log.Println()

	fn := fmt.Sprintf(sp.FILE_NAME, time.Now().Format("2006_01_02"))
	log.Print(fn)

}
