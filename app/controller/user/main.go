package user

import (
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/controller"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/database"
	"net/http"
)

func Show(c echo.Context) error {
	user := model.User{}

	db := database.DB().Preload("House").Preload("House.Point")

	if controller.Authenticated(c, c.Param("userID")) {
		db = db.Preload("Property").Preload("Entries")
	}

	db.Find(&user, c.Param("userID"))

	return c.JSON(http.StatusOK, user)
}
