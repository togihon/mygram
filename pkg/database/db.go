package database

import (
	"fmt"
	"mygram/app/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "togihon"
	dbname   = "mygram"
	db       *gorm.DB
	err      error
)

func Connect() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//create tables
	db.Debug().AutoMigrate(entity.MyGramUser{}, entity.MyGramPhoto{}, entity.MyGramComment{}, entity.MyGramSocialMedia{})
	return db, nil

}
