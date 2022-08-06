package main

import (
	"fmt"
	tomldemo "future-go/toml_demo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println(tomldemo.ConfigCouldReload().Owner.Name)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1) // SIGSEGV
	go func() {
		for {
			<-s
			tomldemo.ReloadConfig()
			log.Println("Reloaded config succ.")

		}
	}()

	http.HandleFunc("/", hello)              // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", tomldemo.ConfigCouldReload().Owner.Name) // 写入到w的是输出到客户端的
}
