package handler

import (
	"log"
	"net/http"
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
	return func(c echo.Context) error {

		input := RegisterRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}

		res, err := uhc.srv.Register(*RequestToCore(input))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}

		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "account succesful created",
		})
	}
}

// Update implements user.UserHandler
func (uhc *userHandlerController) Update() echo.HandlerFunc {
	panic("unimplemented")
}
