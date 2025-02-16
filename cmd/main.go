package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/EkonaldiHutabarat/go-monitoring/internal/controllers"
	"github.com/EkonaldiHutabarat/go-monitoring/internal/database"
	"github.com/EkonaldiHutabarat/go-monitoring/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	//config.LoadEnv()
	wd, _ := os.Getwd()
	fmt.Println("DEBUG: Current working directory:", wd)

	// Cek lokasi file .env yang akan dicari
	envPath := filepath.Join("disini : ", wd, ".env")

	fmt.Println("DEBUG: Looking for .env at:", envPath)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitJWT()
	fmt.Println("Debug: JWT_SECRET = ", os.Getenv("JWT_SECRET"))

	// Init Database
	err = database.InitDB()
	if err != nil {
		log.Fatal("Database connection failed ", err)
	}
	fmt.Println("Connected to PostgresSQL")

	// Setup Router
	r := mux.NewRouter()
	r.HandleFunc("/register", controllers.RegisterUser(database.DB)).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler(database.DB)).Methods("POST")

	// Start Server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
