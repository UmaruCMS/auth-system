package user

import (
	"github.com/UmaruCMS/auth-system/helper"
	"github.com/UmaruCMS/auth-system/model"
)

func RegisterUser(name, email, password string) error {
	user, err := model.NewUser(name, email)
	if err != nil {
		return err
	}
	_, err = model.NewAuthInfo(user.ID, password)
	if err != nil {
		user.Delete(true)
		return err
	}
	return nil
}

func Login(email, password string) (bool, string) {
	user := &model.User{}
	user, err := user.GetByEmail(email)

	if err != nil {
		return false, ""
	}

	authInfo := &model.AuthInfo{}
	authInfo, err = authInfo.GetByUserID(user.ID)

	if err != nil {
		return false, ""
	}

	if authInfo.Password != helper.HashString(password) {
		return false, ""
	}

	tokenString, err := CreateTokenString(authInfo)
	if err != nil {
		return false, ""
	}

	return true, tokenString
}

func GetUserInfo(userID uint) *model.User {
	user := &model.User{}
	user, err := user.GetByID(userID)
	if err != nil {
		return nil
	}
	return user
}
