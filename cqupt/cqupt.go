package cqupt

import (
	"github.com/L2ncE/CQUPT-CourseSelection-Tool/grab"
	"github.com/L2ncE/CQUPT-CourseSelection-Tool/query"
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
