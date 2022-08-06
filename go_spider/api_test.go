package main

import (
	"fmt"
	pictool "future-go/go_spider/pic_tool"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
	"time"
)

// https://worldcosplay.net/member/157692
func TestFapelloGetMomo(t *testing.T) {
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	//fuliba_sort.GetResponse("https://fapello.com/sayo-momo/")
	bobos, err := ioutil.ReadFile("./sayo-momo2.html")
	if err != nil {
		fmt.Println("ReadFile---> ", err)
		return
	}

	re := regexp.MustCompile(`<img src="(.*?)" class=`)
	fixUrls := re.FindAllString(string(bobos), -1)

	for idx, url := range fixUrls {
		tmp := strings.Split(url, "\"")[1]

		if strings.Contains(tmp, "350x600") {
			tmp = tmp[:len(tmp)-11] + "740.jpg"
			fmt.Println(idx, tmp)
			pictool.DownloadPics(tmp, "sayo-momo", "")

			time.Sleep(10 * time.Microsecond)
		}
	}
}

/*
func TestFourBit(t *testing.T) {
	fmt.Printf("-->%04d<--\n", 12) //0012
}

func TestF(t *testing.T) {

	src := "https://fapello.com/content/d/i/disharmonica/%d000/disharmonica_%04d.jpg"

	ch1 := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		go func(i int) {
			for j := 0; j < 1000; j++ {
				url := fmt.Sprintf(src, i, (i-1)*1000+j) //(i/1000)+1, i)

				pictool.DownloadPics(url, "./disharmonica/", "")
				fmt.Println("--->", url)
				time.Sleep(10 * time.Microsecond)
			}
			ch1 <- i
		}(i)
	}
	<-ch1
}

// https://fapello.com/sayo-momo/
func TestFapelloGetPics(t *testing.T) {
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	//fuliba_sort.GetResponse("https://fapello.com/sayo-momo/")
	bobos, err := ioutil.ReadFile("./sayo-momo.html")
	if err != nil {
		fmt.Println("ReadFile---> ", err)
		return
	}

	re := regexp.MustCompile(`<img src="(.*?)" class=`)
	fixUrls := re.FindAllString(string(bobos), -1)

	for idx, url := range fixUrls {
		tmp := strings.Split(url, "\"")[1]

		if strings.Contains(tmp, "300px") {
			tmp = tmp[:len(tmp)-10] + ".jpg"
		}
		fmt.Println(idx, tmp)

		pictool.DownloadPics(tmp, "sayo-momo", "")
		time.Sleep(10 * time.Microsecond)
	}
}


func TestDownloadPics(t *testing.T) {
	picUrl := `https://fapello.com/content/s/a/sayo-momo/1000/sayo-momo_0143.jpg`
	pictool.DownloadPics(picUrl, "./", "")
}

func TestSimple(t *testing.T) {
	sl := []int{23, 56}
	for k, v := range sl {
		log.Printf("--->key:%v,val:%v\n", k, v)
	}

	fn := fmt.Sprintf(sp.FILE_NAME, time.Now().Format("2006_01_02"))
	log.Println(fn)

}

func TestDoc(t *testing.T) {
	var html string = `<tbody id="normalthread_194786">
	<tr>
	<td class="icn">
	<a href="https://www.wnflb99.com/thread-194786-1-1.html" title="有新回复 - 新窗口打开" target="_blank">
	<img src="static/image/common/folder_new.gif">
	</a>
	</td>
	<th class="new">
	<a href="javascript:;" id="content_194786" class="showcontent y" title="更多操作" onclick="CONTENT_TID='194786';CONTENT_ID='normalthread_194786';showMenu({'ctrlid':this.id,'menuid':'content_menu'})"></a>
	 <a href="https://www.wnflb99.com/thread-194786-1-1.html" style="font-weight: bold;color: #2897C5;" onclick="atarget(this)" class="s xst">斯巴达克斯 1-3季+前传</a>
	<img src="static/image/common/agree.gif" align="absmiddle" alt="agree" title="帖子被加分">
	分享 <font color="red">+1</font> [<a href="https://www.wnflb99.com/space-uid-1.html"><font color="#0000ff">admin</font></a> <span title="于 2022-07-17 19:11:38 标为已阅"><a href="https://www.wnflb99.com/forum.php?mod=redirect&amp;goto=findpost&amp;ptid=&amp;pid=3780895"><font color="#ff0000"><b>阅至 1 楼</b></font></a></span>]<a href="https://www.wnflb99.com/forum.php?mod=redirect&amp;tid=194786&amp;goto=lastpost#lastpost" class="xi1">New</a>
	</th>
	<td class="by">
	<cite>
	<a href="https://www.wnflb99.com/space-uid-87517.html" c="1" mid="card_3705">AIPAN</a></cite>
	<em><span><span title="2022-7-17">昨天&nbsp;19:01</span></span></em>
	</td>
	<td class="num"><a href="https://www.wnflb99.com/thread-194786-1-1.html" class="xi2">6</a><em>399</em></td>
	<td class="by">
	<cite><a href="https://www.wnflb99.com/space-username-ohyes200.html" c="1" mid="card_2414">ohyes200</a></cite>
	<em><a href="https://www.wnflb99.com/forum.php?mod=redirect&amp;tid=194786&amp;goto=lastpost#lastpost"><span title="2022-7-17 23:24">昨天&nbsp;23:24</span></a></em>
	</td>
	</tr>
	</tbody>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	time := doc.Find("tr > td:nth-child(3) > em > span > span")
	if time == nil {
		log.Print("nill... ")
	}
	log.Printf("-->%T", time)
	strTime := time.Text()
	//tim, _ := time.Attr("title")
	//tim := time.Text()
	log.Printf("--->[%v] ", strTime)
}

func TestLoliConApi(t *testing.T) {
	url := "https://api.lolicon.app/setu/v2"

	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	dat, _ := ioutil.ReadAll(rsp.Body)
	//log.Print(string(dat))

	var result loli.LoliRsp
	_ = json.Unmarshal(dat, &result)
	fmt.Println(result.Data)
}
*/
