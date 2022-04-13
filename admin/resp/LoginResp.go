package resp

type LoginResp struct {
	AccessToken string `json:"accessToken"`

	RefreshToken string `json:"refreshToken"`
}
