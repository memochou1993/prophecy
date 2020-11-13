package profile

import (
	"github.com/labstack/echo"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/database"
	"gorm.io/gorm/clause"
	"net/http"
)

func Show(c echo.Context) error {
	user := model.User{}

	database.DB().Preload(clause.Associations).Preload("House.Point").Find(&user, 1) // FIXME: use token

	return c.JSON(http.StatusOK, user)
}
