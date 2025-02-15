package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/EkonaldiHutabarat/go-monitoring/internal/database"
	"github.com/EkonaldiHutabarat/go-monitoring/internal/models"
)

func RegisterUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Simpan user ke database
		err := database.InsertUser(db, user)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
	}
}
