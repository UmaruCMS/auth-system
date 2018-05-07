package model

import (
	"github.com/UmaruCMS/auth-system/config"
	"github.com/jinzhu/gorm"
)

type User struct {
	Base
	gorm.Model
	Name  string
	Email string `gorm:"type:varchar(100);unique_index"`
	Role  string `gorm:"size:255"`
}

func NewUser(name, email, role string) (*User, error) {
	db := config.Database
	user := &User{Name: name, Email: email, Role: role}
	if db.NewRecord(*user) {
		db.Create(user)
	}
	return user, nil
}
