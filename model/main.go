package model

import (
	"github.com/UmaruCMS/auth-system/config"
)

func init() {
	db := config.Database
	db.AutoMigrate(&User{}, &AuthInfo{})
}
