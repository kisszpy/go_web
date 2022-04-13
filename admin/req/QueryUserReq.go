package req

type QueryUserReq struct {
	PageReq
	Username string `json:"username"`
}
