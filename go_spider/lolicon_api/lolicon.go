package loliconapi

type UrlList struct {
	Original string `json:"original"`
}

type LoliRspData struct {
	Pid   uint64  `json:"pid"`
	Uid   uint64  `json:"uid"`
	Title string  `json:"title"`
	Urls  UrlList `json:"urls"`
}

type LoliRsp struct {
	ErrCode string        `json:"error"`
	Data    []LoliRspData `json:"data"`
}
