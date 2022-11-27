package cqupt

import (
	"github.com/LgoLgo/cqupt-grabber/grab"
	"github.com/LgoLgo/cqupt-grabber/query"
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
