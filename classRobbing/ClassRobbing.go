package classRobbing

import (
	"encoding/json"
	"github.com/L2ncE/CQUPT-ClassGrabbing/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//SingleRob 仅抢课一次，传递单个load以及cookie
func SingleRob(cookie string, load string) string {
	client := &http.Client{}
	var data = strings.NewReader(load)
	req, err := http.NewRequest("POST", "http://xk1.cqupt.edu.cn/post.php", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Origin", "http://xk1.cqupt.edu.cn")
	req.Header.Set("Referer", "http://xk1.cqupt.edu.cn/yxk.php")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(resp.Body)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var Response model.Response
	err = json.Unmarshal(bodyText, &Response)
	if err != nil {
		log.Fatal(err)
	}
	return Response.Info
}

//LoopRob 循环抢课，支持多个课程同时抢，每次请求停顿0.2秒，防止被ban。
//传入一个cookie和一个load切片
func LoopRob(cookie string, loads []string) {
	for i := 1; ; i++ {
		log.Printf("第%d次抢课开始", i)
		for j, load := range loads {
			j += 1

			//调用SingleRob进行循环抢课

			info := SingleRob(load, cookie)
			if info == "ok" {
				log.Printf("课程%d：%s\n", j, info)
				goto ok
			} else {
				log.Printf("课程%d：%s\n", j, info)
			}
			time.Sleep(250 * time.Millisecond)
		}
		log.Printf("第%d次抢课失败\n\n", i)
		time.Sleep(250 * time.Millisecond)
	}
ok:
	log.Println("抢课成功")
}
