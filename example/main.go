package main

import (
	"github.com/LgoLgo/cqupt-grabber/cqupt"
)

func main() {
	tool := cqupt.New()

	cookie := "your cookie"

	// 可以填入多个关键字进行搜索
	want := []string{
		"这里",
		"可以填入",
		"多个关键字",
	}
	tool.Queryer.Search(cookie, want)
	tool.Queryer.AllForeign(cookie)
	tool.Queryer.AllZiRan(cookie)
	tool.Queryer.AllRenWen(cookie)

	// 支持同时抢多门课程
	loads := []string{
		"这里是第一节课",
		"这里是第二节课",
	}

	tool.Grabber.LoopRob(cookie, loads)
}
