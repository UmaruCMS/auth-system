package user

import (
	"fmt"
	"time"

	"github.com/UmaruCMS/auth-system/model"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	jwt.StandardClaims
	ID uint `json:"id"`
}

const oneWeekTimeDuration int64 = 60 * 60 * 24 * 7

func CreateTokenString(authInfo *model.AuthInfo) (string, error) {
	claims := CustomClaims{
		ID: authInfo.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + oneWeekTimeDuration,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(authInfo.Secret))
}

func ParseTokenString(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// TODO: 将 token secret 的获取过程从 MySQL 转移到 Redis
		if claims, ok := token.Claims.(*CustomClaims); ok {
			authInfo := &model.AuthInfo{}
			authInfo, err := authInfo.GetByUserID(claims.ID)
			if err != nil {
				return nil, err
			}
			return []byte(authInfo.Secret), nil
		}
		return nil, fmt.Errorf("parse error")
	})
}

func VerifyTokenString(tokenString string) bool {
	token, err := ParseTokenString(tokenString)
	if err != nil {
		return false
	}
	return token.Valid
}
