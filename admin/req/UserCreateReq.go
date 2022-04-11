package req

type UserCreateReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Nickname        string `json:"nickname"`
	ConfirmPassword string `json:"confirmPassword"`
	Email           string `json:"email"`
}
type IdReq struct {
	Id int `json:"id"`
}
