package model

import (
	"encoding/json"
	"time"

	"github.com/UmaruCMS/auth-system/config"
)

type TokenInfo struct {
	rawString string
	UserID    uint   `json:"user_id"`
	Secret    string `json:"secret"`
}

const oneWeekNsDuration time.Duration = 1000000000 * 60 * 60 * 24 * 7

func NewTokenInfo(authInfo *AuthInfo, tokenString string) (*TokenInfo, error) {
	tokenInfo := &TokenInfo{
		rawString: tokenString,
		UserID:    authInfo.UserID,
		Secret:    authInfo.Secret,
	}

	infoJSON, err := json.Marshal(tokenInfo)
	if err != nil {
		return nil, err
	}

	rc := config.RedisClient
	err = rc.Set(tokenString, infoJSON, oneWeekNsDuration).Err()
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}

func (tokenInfo *TokenInfo) GetFromRedis(tokenString string) (*TokenInfo, error) {
	rc := config.RedisClient

	tokenJSON, err := rc.Get(tokenString).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(tokenJSON), tokenInfo)
	if err != nil {
		return nil, err
	}

	tokenInfo.rawString = tokenString
	return tokenInfo, nil
}

func (tokenInfo *TokenInfo) Delete() error {
	rc := config.RedisClient
	return rc.Del(tokenInfo.rawString).Err()
}
