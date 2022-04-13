package req

type PageReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Limit    int
	Offset   int
}
