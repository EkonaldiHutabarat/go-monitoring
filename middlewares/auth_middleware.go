package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/EkonaldiHutabarat/go-monitoring/utils"
)

// AuthMiddleware untuk validasi JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari Authorization Header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}

		// Pastikan format header "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]

		// Verifikasi token
		claims, err := utils.VerifyJWT(token)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Simpan informasi user ke context (kalau dibutuhkan)
		r.Header.Set("ID", fmt.Sprintf("%d", claims.ID))
		r.Header.Set("Email", claims.Email)

		// Lanjut ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
