package services

import "simple-shop-api/features/user"

type userServiceCase struct {
	query user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userServiceCase{
		query: ud,
	}
}

// Delete implements user.UserService
func (usc userServiceCase) Delete(token interface{}, userID uint) error {
	panic("unimplemented")
}

// Login implements user.UserService
func (usc userServiceCase) Login(email string, password string) (user.UserCore, error) {
	panic("unimplemented")
}

// Register implements user.UserService
func (usc userServiceCase) Register(newUser user.UserCore) (user.UserCore, error) {
	panic("unimplemented")
}

// Update implements user.UserService
func (usc userServiceCase) Update(token interface{}, updateData user.UserCore) (user.UserCore, error) {
	panic("unimplemented")
}
