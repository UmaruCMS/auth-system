package main

import (
	"fmt"

	"github.com/UmaruCMS/auth-system/config"
	"github.com/UmaruCMS/auth-system/controller/user"
)

func release() {
	config.Release()
}

func main() {
	defer release()
	user.RegisterUser("Lawrence", "lawrence.lee@foxmail.com", "123456")
	_, token := user.Login("lawrence.lee@foxmail.com", "123456")
	fmt.Println(token)
}
