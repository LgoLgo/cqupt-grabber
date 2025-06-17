package query

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/LgoLgo/cqupt-grabber/model"
)

type SmallQueryer struct{}

// 针对于小学期的选课的请求
func smallRequest(class, cookie string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://xk1.cqupt.edu.cn/xxq/search.php?action=getBjKebiao&bj="+class, nil)
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

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	return bodyText
}

// 这个方法针对于小学期选课，阻塞搜索
func (q *SmallQueryer) BlockSearch(cookie string, contents []string, class string) (loads []model.MetaData) {
	for loads == nil {
		loads = q.SimpleSearch(cookie, contents, class)
		time.Sleep(250 * time.Millisecond)
	}
	return
}

// 具体的搜索方法，仅仅会搜索一次
func (q *SmallQueryer) SimpleSearch(cookie string, content []string, class string) (loads []model.MetaData) {
	bodyText := smallRequest(class, cookie)
	var classInfo model.ClassInfos
	err := json.Unmarshal(bodyText, &classInfo)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range classInfo.Data {
		if confirmContain(item, content) {
			loads = append(loads, item)
		}
	}
	return
}
