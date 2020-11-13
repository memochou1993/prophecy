package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/memochou1993/prophecy/app/controller/profile"
	"github.com/memochou1993/prophecy/database"
)

func main() {
	database.Migrate()

	e := echo.New()

	api := e.Group("/api")
	api.GET("/profile", profile.Show)

	e.Logger.Fatal(e.Start(":9000"))
}
