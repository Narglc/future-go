package fuliba_sort

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var Client http.Client

func GetResponse(url string, hdr map[string]string, cks []http.Cookie) *http.Response {
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

	return rsp
}

func SpiderWork(count int) {
	base_url := "https://www.wnflb99.com/forum-2-%d.html" //"https://api.lolicon.app/setu/v2"

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

	data := &[]*FulibaSub{}
	for i := 1; i <= count; i++ {
		log.Println("------> page:", i)
		rsp := GetResponse(fmt.Sprintf(base_url, i), hdr, cks)
		defer rsp.Body.Close()
		if rsp.StatusCode != 200 {
			log.Printf("rsp code:%d\n", rsp.StatusCode)
			continue
		}

		AnalyzeHtml(rsp, data)
	}

	DisplaySortData(data)

	elapsed := time.Since(start)
	log.Printf("Time %s\n", elapsed)
}
