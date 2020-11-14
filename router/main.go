package router

import (
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/controller/house"
	"github.com/memochou1993/prophecy/app/controller/question"
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
	api.POST("/user/houses", house.Store)
	// api.DELETE("/user/houses/:houseID", house.Destroy)
	api.GET("/users/:userID", user.Show)
	api.POST("/houses/:houseID/users/:userID", house.AttachUser)   // FIXME
	api.DELETE("/houses/:houseID/users/:userID", house.DetachUser) // FIXME
	api.GET("/questions", question.Index)

	return router
}
