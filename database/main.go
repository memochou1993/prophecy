package database

import (
	"fmt"
	"github.com/memochou1993/prophecy/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func DB() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate() error {
	return DB().AutoMigrate(
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

func Flush() error {
	return DB().Migrator().DropTable(
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

func Refresh() error {
	if err := Flush(); err != nil {
		log.Fatal(err)
	}

	return Migrate()
}
