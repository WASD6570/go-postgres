package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Getenvs() string {
	db_host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_DATABASE")

	var DSN string = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		db_host, port, user, dbname, pass)

	return DSN
}

func Connection() (*gorm.DB, error) {
	var err error
	Conn, err = gorm.Open(postgres.Open(Getenvs()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return Conn, nil
}
