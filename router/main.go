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
	// api.PATCH("/user/houses/:houseID", house.Update)
	// api.DELETE("/user/houses/:houseID", house.Destroy)
	api.GET("/users/:userID", user.Show)
	api.GET("/questions", question.Index)
	// api.POST("/questions", question.Store)
	// api.PATCH("/questions", question.Update)
	// api.DELETE("/questions", question.Destroy)
	api.POST("/houses/:houseID/users/:userID", house.AttachUser)
	api.DELETE("/houses/:houseID/users/:userID", house.DetachUser)

	return router
}
