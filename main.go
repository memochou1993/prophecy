package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/memochou1993/prophecy/app/controller/house"
	"github.com/memochou1993/prophecy/database"
	"log"
)

func main() {
	e := echo.New()

	if err := database.Migrate(); err != nil {
		log.Fatal(err)
	}

	api := e.Group("/api")
	api.GET("/houses", house.Index)

	e.Logger.Fatal(e.Start(":9000"))
}
