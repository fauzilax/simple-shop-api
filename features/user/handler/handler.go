package handler

import (
	"log"
	"net/http"
	"simple-shop-api/features/user"
	"strconv"

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
	return func(c echo.Context) error {
		token := c.Get("user")
		paramID := c.Param("id")
		employeeID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}
		err = uhc.srv.Delete(token, uint(employeeID))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "data not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success deactivate employee profile",
		})
	}
}

// Login implements user.UserHandler
func (uhc *userHandlerController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		err := c.Bind(input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		res, err := uhc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "account succesful created",
		})
	}
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
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "account succesful created",
		})

	}
}
