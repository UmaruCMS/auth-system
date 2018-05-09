package main

import (
	"github.com/UmaruCMS/auth-system/config"
	"github.com/UmaruCMS/auth-system/http/router"
)

func release() {
	config.Release()
}

func main() {
	defer release()
	// user.RegisterUser("Lawrence", "lawrence.lee@foxmail.com", "123456")
	// _, tokenString := user.Login("lawrence.lee@foxmail.com", "123456")
	// fmt.Println(tokenString)
	// tokenInfo := &model.TokenInfo{}
	// tokenInfo.GetFromRedis(tokenString)
	// tokenInfo.Delete()
	// fmt.Println(user.VerifyTokenString(token))

	r := router.DefaultRouter()
	router.RegisterHandlers(r)
	r.Run(":2333")
}
