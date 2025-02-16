package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error Directory:", err)
	}

	envPath := filepath.Join(rootPath, "..", ".env")

	fmt.Println("Trying to load .env from:", envPath)
	err = godotenv.Overload(envPath)
	if err != nil {
		log.Fatal("Error loading .env file from path:", envPath)
	}

	//fmt.Println(".env file successfully loaded form:", envPath)

}

func GetDBConnection() string {
	// Trim space pada password untuk menghindari karakter tersembunyi
	password := strings.TrimSpace(os.Getenv("DB_PASSWORD"))

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		password,
		os.Getenv("DB_NAME"),
	)

	return dsn
}

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set in .env")
	}
	return secret
}
