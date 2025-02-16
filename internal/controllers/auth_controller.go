package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/EkonaldiHutabarat/go-monitoring/internal/database"
	"github.com/EkonaldiHutabarat/go-monitoring/internal/models"
	"github.com/EkonaldiHutabarat/go-monitoring/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest
		defer r.Body.Close()
		//Decode json req
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println("Error decode request", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		log.Println("Email req:", req.Email)

		// Ambil user dari database berdasarkan email

		user, err := database.GetUserByEmail(db, req.Email)
		if err != nil {
			log.Println("error fetching user", err)
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		log.Println("User ditemukan:", user.Email)

		// Verifikasi password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			log.Println("Password tidak cocok")
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		// Generate JWT token
		token, err := utils.GenerateJWT(user.ID, user.Email)
		if err != nil {
			log.Println("Error generating JWT:", err)
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}
		log.Println("Password cocok, membuat token...")

		// Kirim response dengan token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"responseMessage": "Login Berhasil",
			"token":           token,
		})

	}
}
