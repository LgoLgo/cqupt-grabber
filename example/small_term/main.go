package main

import (
	"github.com/LgoLgo/cqupt-grabber/cqupt"
)

func main() {
	tool := cqupt.NewForSmallTerm()

	str3 := []string{
		"课程关键词1",
		"课程关键词2",
	}
	cookie := "你的 cookie"

	classNo := "你的班级号"

	loads := tool.Queryer.BlockSearch(cookie, str3, classNo)
	tool.Grabber.LoopRob(cookie, loads)
}
