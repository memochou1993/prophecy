package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/controller/profile"
	"github.com/memochou1993/prophecy/app/controller/token"
	"github.com/memochou1993/prophecy/app/middleware"
	"github.com/memochou1993/prophecy/app/request"
	"github.com/memochou1993/prophecy/database"
)

func main() {
	database.Migrate()

	e := request.Validate(echo.New())

	api := e.Group("/api")
	api.POST("/auth/token", token.Issue)

	api.Use(middleware.VerifyToken())
	api.GET("/profile", profile.Show)

	e.Logger.Fatal(e.Start(":9000"))
}
