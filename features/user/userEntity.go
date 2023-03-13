package user

import "github.com/labstack/echo/v4"

type UserCore struct {
	UserID   uint
	Name     string
	Email    string
	Password string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserService interface {
	Register(newUser UserCore) (UserCore, error)
	Login(email string, password string) (UserCore, error)
	Update(token interface{}, updateData UserCore) (UserCore, error)
	Delete(token interface{}, userID uint) error
}

type UserData interface {
	Register(newUser UserCore) (UserCore, error)
	Login(email string) (UserCore, error)
	Update(userID uint, updateData UserCore) (UserCore, error)
	Delete(tokenID uint, userID uint) error
}
