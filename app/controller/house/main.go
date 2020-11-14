package house

import (
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/database"
	"net/http"
)

func AttachUser(c echo.Context) error {
	house := model.House{}
	user := model.User{}

	if result := database.DB().First(&house, c.Param("houseID")); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	if result := database.DB().First(&user, c.Param("userID")); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	if count := database.DB().Model(&house).Where(&user).Association("Users").Count(); count != 0 {
		return c.JSON(http.StatusOK, nil)
	}

	if err := database.DB().Model(&house).Association("Users").Append(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

func DetachUser(c echo.Context) error {
	house := model.House{}
	user := model.User{}

	if result := database.DB().First(&house, c.Param("houseID")); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	if result := database.DB().First(&user, c.Param("userID")); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	if count := database.DB().Model(&house).Where(&user).Association("Users").Count(); count == 0 {
		return c.JSON(http.StatusOK, nil)
	}

	if err := database.DB().Model(&house).Association("Users").Delete(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
