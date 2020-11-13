package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/controller/auth"
	"github.com/memochou1993/prophecy/app/controller/profile"
	"github.com/memochou1993/prophecy/app/middleware"
	"github.com/memochou1993/prophecy/app/request"
	"github.com/memochou1993/prophecy/database"
)

func main() {
	database.Migrate()

	e := request.Validate(echo.New())

	api := e.Group("/api")
	api.POST("/login", auth.Login)

	api.Use(middleware.VerifyToken())
	api.GET("/profile", profile.Show)

	e.Logger.Fatal(e.Start(":9000"))
}
