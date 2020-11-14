package question

import (
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/app/request"
	"github.com/memochou1993/prophecy/database"
	"net/http"
)

func Index(c echo.Context) error {
	question := new(request.Question)

	if err := c.Bind(question); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(question); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var questions []model.Question

	database.
		DB().
		Where(question).
		Preload("Owner").
		Find(&questions)

	return c.JSON(http.StatusOK, questions)
}
