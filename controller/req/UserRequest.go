package req

// LoginRequest 登录请求封装
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
