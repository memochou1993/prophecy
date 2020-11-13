package migration

import (
	"github.com/labstack/echo"
	"github.com/memochou1993/prophecy/app/model"
	"github.com/memochou1993/prophecy/database"
	"log"
)

func Store(echo.Context) error {
	db, err := database.DB()

	if err != nil {
		log.Fatal(err.Error())
	}

	return db.AutoMigrate(
		&model.User{},
		&model.House{},
		&model.Point{},
		&model.Question{},
		&model.Choice{},
		&model.Answer{},
		&model.Participant{},
		&model.Property{},
		&model.Entry{},
	)
}

func Destroy(echo.Context) error {
	db, err := database.DB()

	if err != nil {
		log.Fatal(err.Error())
	}

	return db.Migrator().DropTable(
		&model.User{},
		&model.House{},
		&model.Point{},
		&model.Question{},
		&model.Choice{},
		&model.Answer{},
		&model.Participant{},
		&model.Property{},
		&model.Entry{},
	)
}
