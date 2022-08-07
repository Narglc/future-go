package main

import (
	"flag"
	"fmt"
	tomldemo "future-go/toml_demo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func flagDemo() {
	var user string
	var password string
	var host string
	var port int

	// "u" 指的-u是user参数
	flag.StringVar(&user, "u", "root", "用户名,默认为root")
	flag.StringVar(&password, "p", "", "默认为空")
	flag.StringVar(&host, "h", "localhost", "host参数,默认为localhost")
	flag.IntVar(&port, "port", 999, "端口")

	// 转换,必须调用该方法
	flag.Parse()
	fmt.Println(user, password, host, port)
}

func tomlTest() {
	fmt.Println(tomldemo.ConfigCouldReload().Owner.Name)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGSEGV) //SIGUSR1) // SIGSEGV
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

func main() {
	flagDemo()
	tomlTest()
}
