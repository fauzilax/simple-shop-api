package services

import (
	"errors"
	"simple-shop-api/features/user"
	"simple-shop-api/helper"
)

type userServiceCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userServiceCase{
		qry: ud,
	}
}

// Register implements user.UserService
func (usc *userServiceCase) Register(newUser user.UserCore) (user.UserCore, error) {
	if len(newUser.Password) == 0 {
		return user.UserCore{}, errors.New("password cannot be empty")
	}
	newUser.Password = helper.GeneratePassword(newUser.Password)
	res, err := usc.qry.Register(newUser)
	if err != nil {
		return user.UserCore{}, errors.New("register error")
	}
	return res, nil
}

// Delete implements user.UserService
func (usc *userServiceCase) Delete(token interface{}, userID uint) error {
	panic("unimplemented")
}

// Login implements user.UserService
func (usc *userServiceCase) Login(email string, password string) (user.UserCore, error) {
	panic("unimplemented")
}

// Update implements user.UserService
func (usc *userServiceCase) Update(token interface{}, updateData user.UserCore) (user.UserCore, error) {
	panic("unimplemented")
}
