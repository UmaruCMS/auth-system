package model

import (
	"github.com/UmaruCMS/auth-system/config"
	"github.com/jinzhu/gorm"
)

type User struct {
	Base
	gorm.Model
	Name  string `gorm:"type:varchar(45)"`
	Email string `gorm:"type:varchar(45);unique_index"`
}

func NewUser(name, email string) (*User, error) {
	db := config.Database
	user := &User{Name: name, Email: email}
	db.Create(user)
	return user, nil
}
