package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/LgoLgo/cqupt-grabber/model"
)

type Queryer struct {
}

func request(str, cookie string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://xk1.cqupt.edu.cn/json-data-yxk.php?type=jcts"+str, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	return bodyText
}

func (q *Queryer) AllRenWen(cookie string) {
	bodyText := request("Rw", cookie)
	var classInfo model.ClassInfos
	err := json.Unmarshal(bodyText, &classInfo)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range classInfo.Data {
		var builder strings.Builder
		SRsLimit := strconv.Itoa(item.RsLimit)
		SRwType := strconv.Itoa(item.RwType)

		var strs = []string{"xnxq=", item.Xnxq, "&jxb=", item.Jxb, "&kchb=", item.Kcbh, "&kcmc=", item.Kcmc,
			"&xf=", item.Xf, "&teaname=", item.TeaName, "&rslimit=", SRsLimit, "&rwtype=", SRwType, "&kclb=",
			item.Kclb, "&kchtye=", item.KchType, "&memo=", item.Memo}

		for _, str := range strs {
			builder.WriteString(str)
		}
		loads := builder.String()
		err = ioutil.WriteFile("./output_renwen.txt", []byte(loads), 0666) //写入文件(字节数组)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	return
}

func (q *Queryer) AllZiRan(cookie string) {
	bodyText := request("Zr", cookie)
	var classInfo model.ClassInfos
	err := json.Unmarshal(bodyText, &classInfo)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range classInfo.Data {
		var builder strings.Builder
		SRsLimit := strconv.Itoa(item.RsLimit)
		SRwType := strconv.Itoa(item.RwType)

		var strs = []string{"xnxq=", item.Xnxq, "&jxb=", item.Jxb, "&kchb=", item.Kcbh, "&kcmc=", item.Kcmc, "&xf=",
			item.Xf, "&teaname=", item.TeaName, "&rslimit=", SRsLimit, "&rwtype=", SRwType, "&kclb=", item.Kclb,
			"&kchtye=", item.KchType, "&memo=", item.Memo}

		for _, str := range strs {
			builder.WriteString(str)
		}
		loads := builder.String()
		err = ioutil.WriteFile("./output_ziran.txt", []byte(loads), 0666) //写入文件(字节数组)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	return
}

func (q *Queryer) Search(param, cookie, content string) {
	bodyText := request(param, cookie)
	var classInfo model.ClassInfos
	err := json.Unmarshal(bodyText, &classInfo)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range classInfo.Data {
		if strings.Contains(item.Kcmc, content) {
			fmt.Println("开课学期：", item.Xnxq, "课程名：", item.Kcmc, "学分：", item.Xf, "教师姓名：", item.TeaName)
			var builder strings.Builder
			SRsLimit := strconv.Itoa(item.RsLimit)
			SRwType := strconv.Itoa(item.RwType)

			var strs = []string{"xnxq=", item.Xnxq, "&jxb=", item.Jxb, "&kchb=", item.Kcbh, "&kcmc=", item.Kcmc, "&xf=", item.Xf, "&teaname=", item.TeaName, "&rslimit=", SRsLimit, "&rwtype=", SRwType, "&kclb=", item.Kclb, "&kchtye=", item.KchType, "&memo=", item.Memo}
			for _, str := range strs {
				builder.WriteString(str)
			}
			load := builder.String()
			log.Println(load)
		}
	}
	return
}
