package loggerdemo

import (
	"fmt"
	"log"
	"os"
)

/* SetFlags 支持的参数，如下：
const (
    Ldate         = 1 << iota     // 日期：2022/07/31
    Ltime                         // 时间：16:53:23
    Lmicroseconds                 // 微秒级别时间：16:53:23.468018
    Llongfile                     // 文件全路径名+行号： G:/goProject/06_demo.go:10
    Lshortfile                    // 文件名+行号：06_demo.go:10
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
*/

// 日志简单书到命令行
func log2term() {
	// 使用 SetFlags 函数 修改 Logger 配置
	// 文件全路径名+行号，错误的时间(精确到微秒级别)，错误信息
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	log.Println("输出普通日志") // 每条输出信息都携带了当下时间
	v := "xxxx"
	log.Printf("输出格式化 %s 日志 \n", v)

	// 使用 SetPrefix 对日志的前缀进行设置
	log.SetPrefix("[社区Svr] ")
	// log.Fatalln("输出fatal的日志")
	// log.Panicln("输出panic的日志")
}

// 日志输出到指定日志文件
func log2File() {
	logFile, err := os.OpenFile("./error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败，错误信息:", err)
		return
	}

	// 设置输出文件：使用SetOutput配置日志输出位置
	log.SetOutput(logFile)

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("无前缀到日志输出")
	log.SetPrefix("[丽海不再] ")
	log.Println("有前缀的日志输出")
}

// 使用 New 创建 logger
/* New的原型：
func New(out io.Writer, prefix string, flag int) *Logger
*/
func newLogger() {
	// 实战遇到简单日志，可以直接采用 New 相关写法即可
	logger := log.New(os.Stdout, "[橡皮擦New专用前缀]", log.Lshortfile|log.Ldate)
	logger.Println("普通日志")
}

func TestLogger() {
	log2term()
	log2File()
	newLogger()
}
