package db

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

var DB *pg.DB

func Connect() {
	DB = pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	if DB == nil {
		fmt.Println("Failed to connect to PostgreSQL database")
		os.Exit(1)
	}

	fmt.Println("Connected to PostgreSQL database")
}
