package main

import (
	"fmt"
	"log"
//"strconv"
//"regexp"
	"net/http"
	"io/ioutil"
//"encoding/json"

	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	//"os"
	//"github.com/PuerkitoBio/goquery"
	"strings"
	"bytes"
	"bufio"
	"io"
)


func getPage1(w http.ResponseWriter, r *http.Request) {
	url := "http://query.shenzhentong.com:8080/sztnet/qryCard.do?cardno="
	r.ParseForm()

	//如果不存在卡号
	if len(r.Form["cardno"]) <= 0 {
		return
	}
	url += r.Form["cardno"][0]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	robots, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	sr := strings.NewReader(string(robots))
	tr := transform.NewReader(sr, simplifiedchinese.GB18030.NewEncoder())

	var b bytes.Buffer
	write := bufio.NewWriter(&b)
	io.Copy(write, tr)
	result := b.String()
	fmt.Println(result)
	//result, err := iconv.ConvertString(string(robots), "gbk", "utf-8")
/*
	doc, err := goquery.NewDocument(result)

	doc.Find(".tableact").Each(func(i int, contentSelection *goquery.Selection) {

		info := contentSelection.Find("td").Text()
		fmt.Println(info)
	})

	//写入文件
	if err != nil {
		fp :=  "szt.html"
		err = ioutil.WriteFile(fp, []byte(result), os.ModeAppend)
		if err != nil {
			log.Fatal(err)
		}
	}

	//fmt.Println(result)*/
	fmt.Fprint(w, result)
}

func main() {
	http.HandleFunc("/", getPage1)
	http.ListenAndServe(":8888", nil)
}
