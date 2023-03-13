package handler

import (
	"simple-shop-api/features/user"

	"github.com/labstack/echo/v4"
)

type userHandlerController struct {
	srv user.UserService
}

func New(us user.UserService) user.UserHandler {
	return &userHandlerController{
		srv: us,
	}
}

// Delete implements user.UserHandler
func (uhc *userHandlerController) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// Login implements user.UserHandler
func (uhc *userHandlerController) Login() echo.HandlerFunc {
	panic("unimplemented")
}

// Register implements user.UserHandler
func (uhc *userHandlerController) Register() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements user.UserHandler
func (uhc *userHandlerController) Update() echo.HandlerFunc {
	panic("unimplemented")
}
