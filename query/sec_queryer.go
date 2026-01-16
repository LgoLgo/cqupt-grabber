package query

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/LgoLgo/cqupt-grabber/model"
)

const (
	secRw      = 0
	secZr      = 1
	secForeign = 2
)

var secOptions = []string{"rxrw", "rxzr", "yyxx"}

type SecQueryer struct{}

func secRequest(str, cookie string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://xk1.cqupt.edu.cn/data/json-data.php?type="+str, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Host", "xk1.cqupt.edu.cn")
	req.Header.Set("Referer", "http://xk1.cqupt.edu.cn/xuanke.php")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	return bodyText
}

func (q *SecQueryer) AllRenWen(cookie string) {
	bodyText := secRequest(options[RenWen], cookie)
	loads := prepareFile(bodyText)
	err := os.WriteFile("./output_renwen.txt", []byte(loads), 0o666) // 写入文件(字节数组)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

func (q *SecQueryer) AllZiRan(cookie string) {
	bodyText := secRequest(options[ZiRan], cookie)
	loads := prepareFile(bodyText)
	err := os.WriteFile("./output_ziran.txt", []byte(loads), 0o666) // 写入文件(字节数组)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

func (q *SecQueryer) AllForeign(cookie string) {
	bodyText := secRequest(options[Foreign], cookie)
	loads := prepareFile(bodyText)
	err := os.WriteFile("./output_foreign.txt", []byte(loads), 0o666) // 写入文件(字节数组)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

func (q *SecQueryer) Search(param, cookie, content string) {
	bodyText := secRequest(param, cookie)
	var Resp model.SecResponse
	err := json.Unmarshal(bodyText, &Resp)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range Resp.Data {
		if strings.Contains(item.Kcmc, content) {
			load, err := json.Marshal(item)
			if err != nil {
				log.Println("解析错误")
			}
			log.Println(load)
		}
	}
	return
}

// BlockSearch 阻塞式搜索，直到loads搜索到才会返回
func (q *SecQueryer) BlockSearch(cookie string, contents []string) (loads []model.SecCourseData) {
	for loads == nil {
		loads = q.SimpleSearch(cookie, contents)
		time.Sleep(250 * time.Millisecond)
	}
	return
}

// SimpleSearch 传入需要被搜索的 key 关键字切片，返回 loads
func (q *SecQueryer) SimpleSearch(cookie string, content []string) (loads []model.SecCourseData) {
	for _, option := range secOptions {
		bodyText := secRequest(option, cookie)
		var Resp model.SecResponse
		err := json.Unmarshal(bodyText, &Resp)
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range Resp.Data {
			if confirmContain(fmt.Sprintln(item), content) {
				loads = append(loads, item)
			}
		}
	}
	return
}

// 获取体育课
func (q *SecQueryer) SearchForPE(cookie string, content []string) (loads []model.SecCourseData) {
	MustCourse := secRequest("bj", cookie)
	var MustCourseJson model.SecResponse
	err := json.Unmarshal(MustCourse, &MustCourseJson)
	if err != nil {
		log.Fatal(err)
	}
	for _, course := range MustCourseJson.Data {
		if !strings.Contains(course.Kcmc, "体育") {
			continue
		}
		// 此处获取了所有的体育课程
		PE := secRequest("tyfx&jxb="+course.Jxb, cookie)
		Response := make(map[string]interface{})
		json.Unmarshal(PE, &Response)
		log.Println(Response)
		for _, data := range Response["data"].([]interface{}) {
			mp := data.(map[string]interface{})
			if err != nil {
				log.Fatal(err)
			}
			for _, x := range mp {
				log.Println(x)
			}
			// 确保包含关键词
			if confirmContain(mp["jxbName"].(string), content) ||
				confirmContain(mp["teaName"].(string), content) ||
				confirmContain(mp["subJxb"].(string), content) {
				course.SubJxb = data.(map[string]interface{})["subJxb"].(string)
				loads = append(loads, course)
			}
		}
	}
	return
}
