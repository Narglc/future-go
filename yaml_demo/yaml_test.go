package yamldemo

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"gopkg.in/yaml.v2"
)

func GetSimpleConf() SimpleConf {
	conf := SimpleConf{}

	err := yaml.Unmarshal([]byte(simple_cfg), &conf) // 将读取的 yaml 文件解析为相应的 struct

	if err != nil {
		log.Fatal("yaml.Unmarschal fail,", err.Error())
	}

	return conf
}

func ConstructYamlConf(conf SimpleConf) []byte {
	data, err := yaml.Marshal(&conf)
	if err != nil {
		log.Fatal("yaml.Marschal fail,", err.Error())
	}
	return data
}

func TestYamlSimpleCfg(t *testing.T) {
	conf := GetSimpleConf()

	fmt.Printf("--- simple_cfg:\n%v\n\n", conf)

	ori_data := ConstructYamlConf(conf)
	fmt.Printf("--- conf dump:\n%s\n\n", string(ori_data))

	fmt.Println("----------->end of TestYamlSimpleCfg")
}

func GetConf() Conf {
	var conf Conf
	yamlFile, err := ioutil.ReadFile("config.yaml") // 加载文件
	if err != nil {
		log.Fatal("ReadFile fail,", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &conf) // 将读取的 yaml 文件解析为相应的 struct

	if err != nil {
		log.Fatal("yaml.Unmarschal fail,", err.Error())
	}

	return conf
}

/* 使用go读取yaml文件，并转换为struct结构体 */

func TestYamlLoad(t *testing.T) {
	conf := GetConf()

	fmt.Println(conf.Config.Models[0].Name)
	fmt.Println(conf.Config.Models[0].Schema)

	fmt.Println(conf.Config.Models[1].Name)
	fmt.Println(conf.Config.Models[1].Schema)

	fmt.Println(conf.Config.Acls[0].Model)
	fmt.Println(conf.Config.Acls[0].Role)
	fmt.Println(conf.Config.Acls[1].Operation)
	fmt.Println(conf.Config.Acls[1].Action)

	fmt.Println("----------->end of TestYamlLoad")
}
