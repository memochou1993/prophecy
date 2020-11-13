package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/memochou1993/prophecy/app/controller/token"
	"os"
)

func VerifyToken() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &token.Claims{},
		SigningKey: []byte(os.Getenv("APP_KEY")),
	}

	return middleware.JWTWithConfig(config)
}
