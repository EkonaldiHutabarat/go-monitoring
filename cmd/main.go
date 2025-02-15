// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/EkonaldiHutabarat/go-monitoring/config"
// 	"github.com/EkonaldiHutabarat/go-monitoring/internal/database"
// )

// func main() {
// 	config.LoadEnv()
// 	err := database.InitDB()
// 	if err != nil {
// 		log.Fatal("Database connection failed ", err)
// 	}
// 	fmt.Println("Server is running...")
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EkonaldiHutabarat/go-monitoring/config"
	"github.com/EkonaldiHutabarat/go-monitoring/internal/controllers"
	"github.com/EkonaldiHutabarat/go-monitoring/internal/database"
	"github.com/gorilla/mux"
)

func main() {
	// Load ENV
	config.LoadEnv()

	// Init Database
	err := database.InitDB()
	if err != nil {
		log.Fatal("Database connection failed ", err)
	}
	fmt.Println("Connected to PostgresSQL")

	// Setup Router
	r := mux.NewRouter()
	r.HandleFunc("/register", controllers.RegisterUser(database.DB)).Methods("POST")

	// Start Server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
