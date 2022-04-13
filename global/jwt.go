package global

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var key = []byte("kisszpy")

func GetToken(userId string, issuer string) string {
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	authToken, _ := token.SignedString(key)
	return authToken
}

func Verify(tokenStr string) (*jwt.RegisteredClaims, error) {

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		fmt.Printf("%v \n", claims)
		return key, nil
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
