package user

import (
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/database"
	"net/http"
)

func Show(c echo.Context) error {
	user := model.User{}

	database.
		DB().
		Preload("OwnedHouses").
		Preload("JoinedHouses").
		Preload("OwnedQuestions").
		Preload("JoinedQuestions").
		Preload("Choices").
		Preload("Properties").
		Preload("Properties.Point").
		Preload("Entries").
		Preload("Entries.Point").
		Find(&user, c.Param("userID"))

	return c.JSON(http.StatusOK, user)
}
