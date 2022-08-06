package tomldemo

import (
	"fmt"
	"testing"
)

/* toml 官方库文档
https://github.com/toml-lang/toml
*/

func TestTomlCfg(t *testing.T) {
	fmt.Println(Config().Title)
	fmt.Println(Config().Owner)
	for key, svr := range Config().Servers {
		fmt.Println(key, svr)
	}
}
