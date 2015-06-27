/**
 * shenzhentong.php
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
	//"regexp"
	
	"net/http"
	"io/ioutil"
	//"encoding/json"
	
	iconv "github.com/djimenez/iconv-go"
	
)

func getPage(w http.ResponseWriter, r *http.Request) {
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

	defer resp.Body.Close()
	
	getinfo, _ := ioutil.ReadAll(resp.Body)

	//gbk转码为utf-8
	putinfo := make([]byte, len(getinfo))
	putinfo = putinfo[:]
	iconv.Convert(getinfo, putinfo, "gbk", "utf-8")
	
	result := string(putinfo) //[]byte convter string
	fmt.Println(result)
	fmt.Fprint(w, result)

}
func main() {

	http.HandleFunc("/", getPage)
	http.ListenAndServe(":8888", nil)
}
