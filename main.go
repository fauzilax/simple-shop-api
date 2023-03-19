package main

import (
	"log"
	"simple-shop-api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
