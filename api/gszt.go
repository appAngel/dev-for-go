/**
 * shenzhentong.go
 * 深圳通的API
 * @autuor: Skiychan <dev@skiy.net>
 * @website: www.skiy.net
 * @date: 2015-06-28
 * @readme https://github.com/skiy/dev-for-go/blob/master/docs/shenzhentong.md
 */

/**
链接：http://query.shenzhentong.com:8080/sztnet/qrycard.jsp
接口信息
URL：http://query.shenzhentong.com:8080/sztnet/qryCard.do
     http://query.shenzhentong.com:8080/sztnet/qryCard.do?cardno=328375558
POST方法：cardno:328375558
### 返回字段 json格式
返回值字段 | 字段类型 | 字段说明
----|------|----
card_number   | int     | 卡号
card_balance  | string  | 卡内余额
balance_time  | string  | 余额截止时间
card_validity | string  | 卡有效期
current_time  | string  | 查询时间
*/

package main

import (
	"fmt"
	"log"
	//"strconv"
	//"regexp"
	"net/http"
	"io/ioutil"
	//"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	//"os"
)

/**
 getPage
 通过号码采集页面源码,并分析取出可用数据
*/
func getPage(w http.ResponseWriter, r *http.Request) {
	//url := "http://query.shenzhentong.com:8080/sztnet/qryCard.do?cardno=328375558"
	url := "http://127.0.0.1/328375558.html"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	resbody := mahonia.NewDecoder("gbk").NewReader(resp.Body)
	_, err = ioutil.ReadAll(resbody)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(resbody)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".tableact tr td").Each(func(i int, ss *goquery.Selection) {
		fmt.Print(i)
		fmt.Println(" ", ss.Text())
	})

	fmt.Println("What")
	//fmt.Println(doc)
	//fmt.Fprint(w, doc)
}

func main() {
	http.HandleFunc("/", getPage)
	http.ListenAndServe(":8888", nil)
}
