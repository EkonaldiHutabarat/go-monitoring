package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("Failed to open database: %w", err)

	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("Failed to open database: %w", err)
	}

	DB = db
	//fmt.Println("Connected to PostgresSQL")
	return nil

}
