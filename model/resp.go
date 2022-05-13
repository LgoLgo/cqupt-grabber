package model

type Response struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

type ClassInfos struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data []struct {
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
	} `json:"data"`
}
