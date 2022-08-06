package tomldemo

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

const (
	CFG_FILE = "C:/Users/Narglc/go/src/future-go/toml_demo/config.toml"
)

/* 配置的单例模式 */
var (
	cfg     *tomlConfig
	once    sync.Once
	cfgLock = new(sync.RWMutex)
)

/*
once.Do(func())
使用了sync.Once的Do方法，Do方法当且仅当第一次被调用时才执行函数.
如果once.Do(f)被多次调用，只有第一次调用会执行f，即使f每次调用Do 提供的f值不同
*/
func Config() *tomlConfig {
	once.Do(func() { //
		filePath, err := filepath.Abs(CFG_FILE)
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}

func ConfigCouldReload() *tomlConfig {
	once.Do(ReloadConfig)
	cfgLock.Lock()
	defer cfgLock.Unlock()
	return cfg
}

func ReloadConfig() {
	filePath, err := filepath.Abs(CFG_FILE)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parse toml file once. filePath: %s\n", filePath)

	config := new(tomlConfig)
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		panic(err)
	}
	cfgLock.Lock()
	defer cfgLock.Unlock()
	cfg = config
}
