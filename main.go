package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/memochou1993/prophecy/database"
	"github.com/memochou1993/prophecy/router"
)

func main() {
	database.Migrate()

	e := router.New()

	e.Logger.Fatal(e.Start(":9000"))
}
