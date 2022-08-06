package pictool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func DownloadPics(picUrl, savePath, picName string) {
	rsp, err := http.Get(picUrl)
	if err != nil {
		fmt.Printf("get pics fail:%v\n", err)
		return
	}
	defer rsp.Body.Close()

	bytes, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		fmt.Printf("read fail:%v\n", err)
		return
	}

	var fileName string
	if picName != "" {
		fileName = savePath + picName
	} else {
		tmp := strings.Split(picUrl, "/")
		fileName = savePath + tmp[len(tmp)-1]
	}

	err = ioutil.WriteFile(fileName, bytes, 0666)

	if err != nil {
		fmt.Printf("write fail:%v\n", err)
		return
	}
}
