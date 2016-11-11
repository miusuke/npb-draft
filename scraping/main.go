package main

import (
    "github.com/PuerkitoBio/goquery"
    "fmt"
    "io/ioutil"
    "os"
//    "strings"
//    "github.com/microcosm-cc/bluemonday"
)


func main() {
    for i := 1; i <= 12; i++ {
        saveFile("http://www.baseball-lab.jp/player/"+fmt.Sprint(i), "team"+fmt.Sprint(i))
        saveFile("http://www.baseball-lab.jp/player/pitcher/"+fmt.Sprint(i)+"/2016", "pitch"+fmt.Sprint(i))
        saveFile("http://www.baseball-lab.jp/player/batter/"+fmt.Sprint(i)+"/2016", "bat"+fmt.Sprint(i))
    }
}
/*
func main() {
    doc, err := goquery.NewDocument("http://www.baseball-lab.jp/player/pitcher/2/2016/")
    if err != nil {
        fmt.Print("url scarapping failed")
    }
    doc.Find("a").Each(func(_ int, s *goquery.Selection) {
          url, _ := s.Attr("href")
          fmt.Println(url)
    })
}
*/
/*

対象のページをファイルに保存する
http://www.baseball-lab.jp/1~12
http://www.baseball-lab.jp/player/pitcher/2/2016/
http://www.baseball-lab.jp/player/batter/2/2016/
*/
func saveFile(url string, fileName string) {
    doc, err := goquery.NewDocument(url)
    if err != nil {
        fmt.Print("url scarapping failed")
    }
    res, err := doc.Find("body").Html()
    if err != nil {
        fmt.Print("dom get failed")
    }
    ioutil.WriteFile("./"+fileName+".html", []byte(res), os.ModePerm)
}
/*
func readFile() {
    fileInfos, _ := ioutil.ReadFile("/path/to/goquery.html")
    stringReader := strings.NewReader(string(fileInfos))
    doc, err := goquery.NewDocumentFromReader(stringReader)
    if err != nil {
        fmt.Print("url scarapping failed")
    }
    doc.Find("table > tbody > tr > td.content > span > a").Each(func(_ int, s *goquery.Selection) {
          url, _ := s.Attr("href")
          fmt.Println(url)
    })
}

func dataSanitize() {
    p := bluemonday.NewPolicy()

    p.AllowElements("li").AllowElements("ul")

    html := p.Sanitize(
        `<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E5.AD.97.E6.BA.90"><span class="tocnumber">1.1</span> <span class="toctext">字源</span></a></li>
<li class="toclevel-2 tocsection-3"><a href="#.E9.9F.B3"><span class="tocnumber">1.2</span> <span class="toctext">音</span></a></li>
<li class="toclevel-2 tocsection-4"><a href="#.E6.84.8F.E7.BE.A9"><span class="tocnumber">1.3</span> <span class="toctext">意義</span></a>
</ul>
`,
    )

    fmt.Println(html)
}
*/
