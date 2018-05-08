package main

import (
	"github.com/UmaruCMS/auth-system/config"
)

func release() {
	config.Release()
}

func main() {
	defer release()
	// user.RegisterUser("Lawrence", "lawrence.lee@foxmail.com", "123456")
	// _, tokenString := user.Login("lawrence.lee@foxmail.com", "123456")
	// fmt.Println(tokenString)
	// tokenInfo, err := json.Marshal(model.TokenInfo{})
	// if err != nil {
	// 	panic(err)
	// }
	// err = config.RedisClient.Set(tokenString, tokenInfo, oneWeekTimeDuration).Err()
	// fmt.Println(err)
	// fmt.Println(config.RedisClient.Get(tokenString).Result())
	// fmt.Println(user.VerifyTokenString(token))
	// fmt.Println(config.RedisClient.Ping().Result())
}
