package model

import "github.com/jinzhu/gorm"

type AuthInfo struct {
	Base
	gorm.Model
	UserID   uint   `gorm:"AUTO_INCREMENT"`
	Salt     string `gorm:"type:varchar(45)"`
	PassCode string `gorm:"type:varchar(45)"`
}
