package data

import (
	"simple-shop-api/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return userQuery{
		db: db,
	}
}

// Register implements user.UserData
func (uq userQuery) Register(newUser user.UserCore) (user.UserCore, error) {
	panic("unimplemented")
}

// Login implements user.UserData
func (uq userQuery) Login(email string) (user.UserCore, error) {
	panic("unimplemented")
}

// Update implements user.UserData
func (uq userQuery) Update(userID uint, updateData user.UserCore) (user.UserCore, error) {
	panic("unimplemented")
}

// Delete implements user.UserData
func (uq userQuery) Delete(tokenID uint, userID uint) error {
	panic("unimplemented")
}
