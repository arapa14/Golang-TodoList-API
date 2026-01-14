package main

import (
	"TODO-LIST-API/infrastructure/database"
	"TODO-LIST-API/internal/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// Global variable
var db *sql.DB

func main() {
	// Load .ENV file && Config
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed load .ENV file:", err)
	}
	config := config.Load()

	// Try connect to database
	var err error
	db, err = database.NewPostgres(config.DB)
	if err != nil {
		log.Fatal("Failed connect to database:", err)
	}

	// Route
	http.HandleFunc("/api/v1/todos", HandleTodo)

	// Server
	fmt.Println("Server running at port: 8080")
	http.ListenAndServe(":8080", nil)
}

// Todo handler
func HandleTodo(w http.ResponseWriter, r *http.Request) {

}