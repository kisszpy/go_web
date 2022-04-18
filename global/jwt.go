package global

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GetToken(userId string) string {
	claims := jwt.RegisteredClaims{
		Issuer:    CONF.Jwt.Issuer,
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	authToken, _ := token.SignedString([]byte(CONF.Jwt.Key))
	return authToken
}

func Verify(tokenStr string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(CONF.Jwt.Key), nil
	})
	if token.Valid {
		return claims, nil
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return claims, errors.New("token 过期了")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return claims, errors.New("token 不合法")
	} else {
		return claims, errors.New("错误的token")
	}
}
