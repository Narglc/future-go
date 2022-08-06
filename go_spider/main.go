package main

import (
	fs "future-go/go_spider/fuliba_sort"
	"log"
	"os"
)

func main() {
	// TODO: 发帖,学习写个脚本,排序吧里的热帖
	log.SetFlags(log.LstdFlags)
	log.Print("args= ", len(os.Args))
	var count int = 1
	/*
		if len(os.Args) == 2 {
			fs.SpiderWork(count)
		}
	*/
	fs.SpiderWork(count)

}
