package user

import (
	"time"

	"github.com/UmaruCMS/auth-system/model"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	jwt.StandardClaims
	ID uint `json:"id"`
}

const oneWeekSecDuration int64 = 60 * 60 * 24 * 7

func CreateTokenString(authInfo *model.AuthInfo) (string, error) {
	claims := CustomClaims{
		ID: authInfo.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + oneWeekSecDuration,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(authInfo.Secret))
	if err != nil {
		return "", err
	}
	_, err = model.NewTokenInfo(authInfo, tokenString)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseTokenString(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		tokenInfo := &model.TokenInfo{}
		tokenInfo, err := tokenInfo.GetFromRedis(tokenString)
		if err != nil {
			return nil, err
		}
		return []byte(tokenInfo.Secret), nil
	})
}

func VerifyTokenString(tokenString string) bool {
	token, err := ParseTokenString(tokenString)
	if err != nil {
		return false
	}
	return token.Valid
}
