package question

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/controller/token"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/database"
	"net/http"
)

func Index(c echo.Context) error {
	house := model.House{}

	if result := database.DB().First(&house, c.Param("houseID")); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	var questions []model.Question

	database.
		DB().
		Where("house_id = ?", house.ID).
		Preload("Owner").
		Find(&questions)

	return c.JSON(http.StatusOK, questions)
}

func Store(c echo.Context) error {
	house := model.House{}

	if result := database.DB().First(&house, c.Param("houseID")); result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	userID := c.Get("user").(*jwt.Token).Claims.(*token.Claims).UserID

	if !isMember(userID, house.ID) {
		return c.JSON(http.StatusBadRequest, nil)
	}

	question := model.Question{
		OwnerID: userID,
	}

	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(question); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := database.DB().Model(&house).Association("Questions").Append(&question); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

func isMember(userID uint, houseID uint) bool {
	var count int64

	database.
		DB().
		Model(&model.Member{}).
		Where("user_id", userID).
		Where("house_id", houseID).
		Count(&count)

	return count == 1
}
