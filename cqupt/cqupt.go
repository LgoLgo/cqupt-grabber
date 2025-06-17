package cqupt

import (
	"github.com/LgoLgo/cqupt-grabber/grab"
	"github.com/LgoLgo/cqupt-grabber/query"
)

type Engine struct {
	Queryer query.Queryer
	Grabber grab.Grabber
}

type SmallEngine struct {
	Queryer query.SmallQueryer
	Grabber grab.SmallGrabber
}

// 正常选课时，使用这个实例
func New() *Engine {
	h := &Engine{
		Queryer: query.Queryer{},
		Grabber: grab.Grabber{},
	}
	return h
}

// 小学期选课使用这个获取实例
func NewForSmallTerm() *SmallEngine {
	h := &SmallEngine{
		Queryer: query.SmallQueryer{},
		Grabber: grab.SmallGrabber{},
	}
	return h
}
