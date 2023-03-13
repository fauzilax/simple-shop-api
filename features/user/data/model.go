package data

import (
	"simple-shop-api/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func DataToCore(data User) user.UserCore {
	return user.UserCore{
		UserID:   data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}

func CoreToData(core user.UserCore) User {
	return User{
		Model:    gorm.Model{ID: core.UserID},
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
	}
}
