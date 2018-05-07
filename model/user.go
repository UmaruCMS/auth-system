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
	err := db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) GetByEmail(email string) (*User, error) {
	db := config.Database
	if err := db.Where("email = ?", email).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) Delete(permanently bool) {
	db := config.Database
	if permanently {
		db.Delete(user)
	} else {
		db.Unscoped().Delete(user)
	}
}
