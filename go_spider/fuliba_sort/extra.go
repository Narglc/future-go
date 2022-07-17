package fuliba_sort

import (
	"regexp"
	"strconv"
)

func GetUsefulDat(data string) (lins []string) {
	reg := regexp.MustCompile(`<tbody id="normalthread_.*?">(.*?)</tbody>`)
	ctxs := reg.FindAllString(data, 100)
	return ctxs
}

func MutliCoroSpider() {
	base_url := "http://www.baidu.com"
	var goch = make(chan int)

	for i := 0; i < 6; i++ {
		url := base_url + strconv.Itoa(i)
		go func() {
			_ = GetResponse(url, nil, nil)
		}()
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
