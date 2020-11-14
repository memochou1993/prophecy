package router

import (
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/controller/token"
	"github.com/memochou1993/prophecy/app/controller/user"
	"github.com/memochou1993/prophecy/app/middleware"
	"github.com/memochou1993/prophecy/app/request"
)

func New() *echo.Echo {
	router := request.Validate(echo.New())

	api := router.Group("/api")
	api.POST("/auth/token", token.Issue)

	api.Use(middleware.VerifyToken())
	api.GET("/users/:userID", user.Show)
	// api.GET("/users/:user/questions", nil) // TODO

	return router
}
