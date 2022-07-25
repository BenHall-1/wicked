package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Database *sqlx.DB

func OpenDatabase() {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"),
	)
	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		fmt.Println("Failed to open database: ", err.Error())
		os.Exit(1)
	}

	Database = db
}
