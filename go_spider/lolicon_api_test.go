package main

import (
	"fmt"
	sp "future-go/go_spider/fuliba_sort"
	"log"
	"testing"
)

func TestLoliConApi(t *testing.T) {
	url := "https://api.lolicon.app/setu/v2"

	code, ctx := sp.Spider(url, nil, nil)

	log.Print(code, ctx)
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
}
