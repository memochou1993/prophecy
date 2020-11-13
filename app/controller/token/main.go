package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/app/request"
	"github.com/memochou1993/prophecy/database"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func Issue(c echo.Context) error {
	credentials := new(request.Credentials)

	if err := c.Bind(credentials); err != nil {
		return echo.ErrInternalServerError
	}

	if err := c.Validate(credentials); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	user := model.User{}

	result := database.DB().Where(&model.User{Email: credentials.Email}).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.ErrUnauthorized
	}

	if user.Password != credentials.Password {
		return echo.ErrUnauthorized
	}

	claims := &Claims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("APP_KEY")))

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
