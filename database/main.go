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
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func Migrate() {
	err := DB().AutoMigrate(
		&model.User{},
		&model.House{},
		&model.Member{},
		&model.Question{},
		&model.Participant{},
		&model.Choice{},
		&model.Answer{},
		&model.Point{},
		&model.Property{},
		&model.Entry{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func Flush() {
	err := DB().Migrator().DropTable(
		&model.User{},
		&model.House{},
		&model.Member{},
		&model.Question{},
		&model.Participant{},
		&model.Choice{},
		&model.Answer{},
		&model.Point{},
		&model.Property{},
		&model.Entry{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
