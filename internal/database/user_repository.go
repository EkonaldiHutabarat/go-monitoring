package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/EkonaldiHutabarat/go-monitoring/internal/models"
)

func InsertUser(db *sql.DB, user models.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1,$2,$3)`

	_, err := db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		log.Println("Error inserting user", err)
		return errors.New("Failed to create user")
	}

	fmt.Println("Satu user berhasil register", user.Name)
	return nil
}
