package cqupt

import (
	"github.com/LgoLgo/Lgo-cqupt-grabber/grab"
	"github.com/LgoLgo/Lgo-cqupt-grabber/query"
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
