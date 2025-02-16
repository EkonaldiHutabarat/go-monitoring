package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/EkonaldiHutabarat/go-monitoring/internal/database"
	"github.com/EkonaldiHutabarat/go-monitoring/internal/models"
	"github.com/EkonaldiHutabarat/go-monitoring/utils"
)

func RegisterUser(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		//Hash password sebelum disimpan kedatabase
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Failed to hash pasword", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword //simpan password yg sudah di hash

		// Simpan user ke database
		err = database.InsertUser(db, user)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"responseMessage": "register successfully",
		})
	}
}
