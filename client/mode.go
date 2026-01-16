package client

type Mode string

const (
	ModeCommon    Mode = "普通选课 New()"
	ModeSecXk     Mode = "二次选课 NewForSecXk()"
	ModeSmallTerm Mode = "小学期 NewForSmallTerm()"
)
