package house

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/controller/token"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/database"
	"net/http"
)

func Store(c echo.Context) error {
	user := model.User{}

	if result := database.DB().First(&user, c.Get("user").(*jwt.Token).Claims.(*token.Claims).UserID); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	house := model.House{}

	if err := c.Bind(&house); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(house); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var count int64

	database.DB().Model(&house).Where(&house).Count(&count)

	if count != 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := database.DB().Model(&user).Association("OwnedHouses").Append(&house); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

// func Update(c echo .Context) error {
// 	// TODO
// }

func Destroy(c echo.Context) error {
	user := model.User{}

	if result := database.DB().First(&user, c.Get("user").(*jwt.Token).Claims.(*token.Claims).UserID); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	house := model.House{
		OwnerID: user.ID,
	}

	if result := database.DB().First(&house, c.Param("houseID")); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	database.DB().Model(&house).Delete(&house)

	return c.JSON(http.StatusNoContent, nil)
}

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
