package main

import (
	"github.com/UmaruCMS/auth-system/config"
	"github.com/UmaruCMS/auth-system/model"
)

func release() {
	config.Release()
}

func main() {
	defer release()
	model.NewUser("Lawrence", "lawrence.lee@foxmail.com", "ADMIN")
}
