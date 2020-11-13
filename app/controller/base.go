package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/memochou1993/prophecy/app/controller/token"
	"strconv"
)

func Authenticated(c echo.Context, userID string) bool {
	return strconv.Itoa(int(c.Get("user").(*jwt.Token).Claims.(*token.Claims).UserID)) == userID
}
