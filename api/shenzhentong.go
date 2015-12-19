/**
 * shenzhentong.php
 * 深圳通的API
 * Author: Skiychan 
 * Email : dev@skiy.net
 * Link  : www.zzzzy.com
 * Create: 2015-06-28
 * desc  : https://github.com/skiy/dev/blob/master/docs/shenzhentong.md
 */

/**
链接：http://query.shenzhentong.com:8080/sztnet/qrycard.jsp

接口信息
URL：http://query.shenzhentong.com:8080/sztnet/qryCard.do?cardno=328375558

GET方法：   
参数：cardno
值  ：int

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
    "github.com/PuerkitoBio/goquery"

    "fmt"
    "net/http"
    "github.com/axgle/mahonia"
    "regexp"
    "time"
    "encoding/json"
    "strconv"
)

type Szt struct {
    CardNumber int `json:"card_number"`
    CardBalance string `json:"card_balance"`
    BalanceTime string `json:"balance_time"`
    CardValidity string `json:"card_validity"`
    CurrentTime string `json:"current_time"`
}

func getPage(w http.ResponseWriter, r *http.Request) {
    url := "http://query.shenzhentong.com:8080/sztnet/qryCard.do?cardno="
    r.ParseForm()

    //如果不存在卡号
    if len(r.Form["cardno"]) <= 0 {
        return
    }
    url += r.Form["cardno"][0]

    var doc *goquery.Document
    var err error
    enc := mahonia.NewDecoder("gbk")
    if doc, err = goquery.NewDocument(url); err != nil {
        panic(err.Error())
    }

    var szt Szt
    var val string
    doc.Find(".tableact tr td").Each(func(i int, ss *goquery.Selection) {
        val = enc.ConvertString(ss.Text())

        switch i {
        case 1:
            szt.CardNumber, _ = strconv.Atoi(val)
        case 2:
            preg := `2[0-9-\s:]*`
            re, _ := regexp.Compile(preg)
            szt.BalanceTime = re.FindString(val)
        case 3:
            szt.CardBalance = val
        case 5:
            szt.CardValidity = val
        }
    })

    szt.CurrentTime = time.Now().Format("2006-01-02 15:04:05")
    res, _ := json.Marshal(szt)

    fmt.Fprintf(w, string(res))  //显示到网页
    fmt.Println(string(res))  //控制台显示
}

func main() {
    http.HandleFunc("/", getPage)
    http.ListenAndServe(":8888", nil)
}
