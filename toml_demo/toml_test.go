package tomldemo

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

/* toml 官方库文档
https://github.com/toml-lang/toml

参考文章: https://www.cnblogs.com/CraryPrimitiveMan/p/7928647.html
- 配置的单例模式
- 配置的更新
*/

// 配置的单例模式
func TestTomlCfg(t *testing.T) {
	fmt.Println(Config().Title)
	fmt.Println(Config().Owner)
	for key, svr := range Config().Servers {
		fmt.Println(key, svr)
	}
}

// 运行中更新配置
/*
启动server之后，可以通过如下shell命令更新配置:	kill -SIGUSR1 6666
其中的6666是go server的进程号。执行这条命令之后，会向go server发送syscall.SIGUSR1的信号，从而触发更新配置的动作。
ps. windows 不支持信号 SIGUSR1

# windows下查看进程信息:    tasklist |grep main.exe
# 查看tasklist帮助信息:		tasklist /?
# 杀死进程:					tskill [PID]

*/
func TestTomlUpdateCfg(t *testing.T) {
	fmt.Println(ConfigCouldReload().Owner.Name)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			ReloadConfig()
			log.Println("Reloaded config succ.")

		}
	}()

	http.HandleFunc("/", hello)              // 设置访问的路由
	err := http.ListenAndServe(":9091", nil) // 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", ConfigCouldReload().Owner.Name) // 写入到w的是输出到客户端的
}
