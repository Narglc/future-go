package tools

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/briandowns/spinner"
)

// 使用原生 Fsprintf 打印进度条
func TestProgressBar(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Fprintf(os.Stdout, "result is %d\r", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("over")
}

// 使用第三方库打印进度条
// https://github.com/briandowns/spinner
func TestSpinner(t *testing.T) {
	for i := 0; i < 43; i++ {
		s := spinner.New(spinner.CharSets[i], 100*time.Millisecond)
		s.Color("bgBlack", "bold", "fgRed")
		s.Start()
		time.Sleep(2 * time.Second)
		s.Stop()
	}

}
