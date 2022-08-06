package tools

import (
	"flag"
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
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
