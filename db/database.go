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
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	})

	fmt.Println(os.Getenv("DB_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	if DB == nil {
		fmt.Println("Failed to connect to PostgreSQL database")
		os.Exit(1)
	}

	fmt.Println("Connected to PostgreSQL database")
}
