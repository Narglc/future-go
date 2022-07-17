package fuliba_sort

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func AnalyzeHtml(rsp *http.Response, data *[]*FulibaSub) {
	// 生成一个goquery的doc
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ctx, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(ctx))

	// goquery用的最多的是Find函数，类似于jquery的$()，可以选择dom结构
	doc.Find("tbody[id]").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Attr("id"))
		title := s.Find("tr > th > a.s.xst").Text()
		url, _ := s.Find("tr > th > a.s.xst").Attr("href")
		zan_str := s.Find(`tbody > tr > th > font[color="red"]`).Text()
		var like int = 0
		if zan_str != "" {
			like, _ = strconv.Atoi(zan_str[1:])
		}

		reply_num, _ := strconv.Atoi(s.Find("td.num > a").Text())
		read_num, _ := strconv.Atoi(s.Find("td.num > em").Text())

		//fmt.Printf("\n---->title:%v\n like[%d] num:[%d/%d]\n url:%s", title, like, reply_num, read_num, url)

		if title == "" {
			return
		}
		*data = append(*data, &FulibaSub{
			Title: title,
			Like:  like,
			Reply: reply_num,
			Read:  read_num,
			Url:   url,
		})

	})

}

func DisplaySortData(data *[]*FulibaSub) {
	sort.Sort(FulibaSubs(*data))

	fmt.Printf("%s\t[%s:%s:%s]\t %s[%s]\n", "Top", "Read", "Reply", "Like", "Title", "Url")
	for k, v := range *data {
		fmt.Printf("%d\t[%d:%d:%d]\t [%s]\t %s\n", k, v.Read, v.Reply, v.Like, v.Url, v.Title)
	}
}
