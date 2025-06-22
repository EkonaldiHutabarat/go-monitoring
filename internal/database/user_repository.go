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

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	log.Println("Mencari user dengan email:", email)
	var user models.User
	query := "SELECT id, name, email, password FROM users WHERE email = $1 LIMIT 1"
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			//return nil, errors.New("User not found")
			fmt.Println("error : ", err)
			log.Printf("User tidak ditemukan dengan email: %s", email)

		}
		log.Println("Error saat menjalankan query:", err)
		return nil, err
	}
	return &user, nil
}
