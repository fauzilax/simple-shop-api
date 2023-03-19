package data

import (
	"errors"
	"log"
	"simple-shop-api/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

// Register implements user.UserData
func (uq *userQuery) Register(newUser user.UserCore) (user.UserCore, error) {
	data := CoreToData(newUser)
	err := uq.db.Create(newUser).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.UserCore{}, errors.New("create account fail, problem with server")
	}
	return DataToCore(data), nil
}

// Delete implements user.UserData
func (uq *userQuery) Delete(tokenID uint, userID uint) error {
	panic("unimplemented")
}

// Login implements user.UserData
func (uq *userQuery) Login(email string) (user.UserCore, error) {
	panic("unimplemented")
}

// Update implements user.UserData
func (uq *userQuery) Update(userID uint, updateData user.UserCore) (user.UserCore, error) {
	panic("unimplemented")
}
