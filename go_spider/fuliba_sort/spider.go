package fuliba_sort

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var Client http.Client

func Spider(url string, hdr map[string]string, cks []http.Cookie) (int, string) {
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil) //或直接使用htt.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	//添加cookie
	for _, v := range cks {
		req.AddCookie(&v)
	}

	//添加header
	for k, v := range hdr {
		req.Header.Add(k, v)
	}

	rsp, err := Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	content, _ := ioutil.ReadAll(rsp.Body)

	return rsp.StatusCode, string(content)
}

func GetUsefulDat(data string) (lins []string) {
	reg := regexp.MustCompile(`<tbody id="normalthread_.*?">(.*?)</tbody>`)
	ctxs := reg.FindAllString(data, 100)
	return ctxs
}

func Save2File(fname string, ctx string) {
	var file *os.File
	_, err := os.Lstat(fname)
	if err != nil {
		file, err = os.Create(fname)
	} else {
		file, err = os.OpenFile(fname, os.O_APPEND, 7)
		if err != nil {
			log.Fatal(err)
		}
	}
	file.Write([]byte(ctx))
}

func MutliCoroSpider() {
	base_url := "http://www.baidu.com"
	var goch = make(chan int)

	for i := 0; i < 6; i++ {
		url := base_url + strconv.Itoa(i)
		go Spider(url, nil, nil)
	}
	for i := 0; i < 6; i++ {
		<-goch // 记得在 Spider结束时将值写入通道，否则会一直阻塞
	}
}

/*

sync.WaitGroup

var wg sync.WaitGroup
wg.Add(10)

for i:=range []int{1,2,3}{
	go func(i int){
		defer wg.Done()
		Spider(url,i)
	}(i)
}

wg.Wait()


*/
