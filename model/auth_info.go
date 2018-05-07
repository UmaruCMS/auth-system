package model

import (
	"github.com/UmaruCMS/auth-system/config"

	"github.com/UmaruCMS/auth-system/helper"
	"github.com/jinzhu/gorm"
)

type AuthInfo struct {
	Base
	gorm.Model
	UserID   uint   `gorm:"unique_key"`
	Salt     string `gorm:"type:varchar(10)"`
	Password string `gorm:"type:varchar(45)"`
}

func NewAuthInfo(userID uint, password string) (*AuthInfo, error) {
	db := config.Database
	salt := helper.GetRandomString(10)
	authInfo := &AuthInfo{UserID: userID, Salt: salt, Password: helper.HashString(password)}
	err := db.Create(authInfo).Error
	if err != nil {
		return nil, err
	}
	return authInfo, nil
}

func (authInfo *AuthInfo) GetByUserID(userID uint) (*AuthInfo, error) {
	db := config.Database
	if err := db.Where("user_id = ?", userID).Find(authInfo).Error; err != nil {
		return nil, err
	}
	return authInfo, nil
}
