package main

import (
	"log"
	"simple-shop-api/config"
	usrData "simple-shop-api/features/user/data"
	usrHandler "simple-shop-api/features/user/handler"
	usrService "simple-shop-api/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	uData := usrData.New(db)
	uSrv := usrService.New(uData)
	uHdl := usrHandler.New(uSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.POST("/register", uHdl.Register())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
