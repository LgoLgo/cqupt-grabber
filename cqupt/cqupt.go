package cqupt

import (
	"github.com/L2ncE/cqupt-course-tool/grab"
	"github.com/L2ncE/cqupt-course-tool/query"
)

type Engine struct {
	Queryer query.Queryer
	Grabber grab.Grabber
}

func New() *Engine {
	h := &Engine{
		Queryer: query.Queryer{},
		Grabber: grab.Grabber{},
	}
	return h
}
