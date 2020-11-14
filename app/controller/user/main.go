package user

import (
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/database"
	"net/http"
)

func Show(c echo.Context) error {
	user := model.User{}

	result := database.
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
		First(&user, c.Param("userID"))

	if result.RowsAffected == 0 {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, user)
}
