package main

import (
	"github.com/memochou1993/quiz/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/quiz"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// TODO: should be removed
	err = db.Migrator().DropTable(
		&model.User{},
		&model.House{},
		&model.Game{},
		&model.Question{},
		&model.Choice{},
		&model.Answer{},
		&model.Participant{},
		&model.Entry{},
	)

	err = db.AutoMigrate(
		&model.User{},
		&model.House{},
		&model.Game{},
		&model.Question{},
		&model.Choice{},
		&model.Answer{},
		&model.Participant{},
		&model.Entry{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
