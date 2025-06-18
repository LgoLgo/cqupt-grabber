package model

type Response struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

type ClassInfos struct {
	Code int        `json:"code"`
	Info string     `json:"info"`
	Data []MetaData `json:"data"`
}

type MetaData struct {
	Xnxq    string `json:"xnxq"`
	Jxb     string `json:"jxb"`
	Kcbh    string `json:"kcbh"`
	Kcmc    string `json:"kcmc"`
	Xf      string `json:"xf"`
	TeaName string `json:"teaName"`
	RsLimit int    `json:"rsLimit"`
	RwType  int    `json:"rwType"`
	Kclb    string `json:"kclb"`
	KchType string `json:"kchType"`
	Memo    string `json:"memo"`

	// 以下是小学期独有属性
	KcInfo string `json:"kcInfo"` // 课程信息，小学期独有
	Tea    string `json:"tea"`    // 这部分是小学期的老师信息
	SkInfo string `json:"skInfo"` // 小学期特有属性，上课信息
}

type SmallRequest struct {
	Action string `json:"action"`
	Jxb    string `json:"jxb"`
	Kclb   string `json:"kclb"`
	Kcbh   string `json:"kcbh"`
}
