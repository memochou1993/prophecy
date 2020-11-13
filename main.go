package main

import (
	"github.com/labstack/echo"
	"github.com/memochou1993/prophecy/app/controller/house"
	"github.com/memochou1993/prophecy/app/controller/migration"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	e := echo.New()

	// FIXME: should group with api prefix
	e.GET("/api/houses/:id", house.Index)

	if os.Getenv("APP_ENV") == "local" {
		e.POST("/api/migrations", migration.Store)
		e.DELETE("/api/migrations", migration.Destroy)
	}

	e.Logger.Fatal(e.Start(":9000"))
}
