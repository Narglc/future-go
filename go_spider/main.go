package main

import (
	"fmt"
	sp "future-go/go_spider/fuliba_sort"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetHtmlContent() {
	url := "https://www.wnflb99.com/forum-2-1.html" //"https://api.lolicon.app/setu/v2"

	start := time.Now()

	hdr := map[string]string{
		"referer": "https://www.wnflb99.com/forum.php",
	}

	cks := []http.Cookie{
		http.Cookie{
			Name:  "S5r8_2132_saltkey",
			Value: "Xmw2110g",
		},
		http.Cookie{
			Name:  "S5r8_2132_auth",
			Value: "5120i4nYO5uZWldiuRIiTcDxUNzkLKEc74w9oT1otlehfudkorAhGJJ7I3LhWLM1S9otzF684pD%2BWUBZe1jXJszYbA",
		},
	}

	_, ctx := sp.Spider(url, hdr, cks)
	sp.Save2File("./fuliba.html", ctx)

	elapsed := time.Since(start)
	log.Printf("Time %s\n", elapsed)
}

type FulibaSub struct {
	Title string
	Like  int
	Reply int
	Read  int
	Url   string
}

type FulibaSubs []*FulibaSub

func (f FulibaSubs) Len() int {
	return len(f)
}

func (f FulibaSubs) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f FulibaSubs) Less(i, j int) bool {
	return f[i].Read > f[j].Read
}

func AnalyzeHtml() {
	f, err := os.Open("./fuliba.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 生成一个goquery的doc
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	data := []*FulibaSub{}

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

		data = append(data, &FulibaSub{
			Title: title,
			Like:  like,
			Reply: reply_num,
			Read:  read_num,
			Url:   url,
		})

	})

	sort.Sort(FulibaSubs(data))

	fmt.Printf("%s\t[%s:%s:%s]\t %s[%s]\n", "Top", "Read", "Reply", "Like", "Title", "Url")
	for k, v := range data {
		fmt.Printf("%d\t[%d:%d:%d]\t %s[%s]\n", k, v.Read, v.Reply, v.Like, v.Title, v.Url)
	}

}

func main() {
	AnalyzeHtml()
}
