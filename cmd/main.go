package main

import (
	"TODO-LIST-API/infrastructure/database"
	"TODO-LIST-API/internal/config"
	"TODO-LIST-API/internal/shared"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// Data table 1: Todo
type Todo struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Is_Completed bool   `json:"is_completed"`
	Due_Date     string `json:"due_date"`
	Priority     int    `json:"priority"`
	Created_At   string `json:"created_at"`
	Updated_At   string `json:"updated_at"`
	Deleted_At   string `json:"deleted_at"`
}

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
	switch r.Method {
	case http.MethodGet:
		TodoGetUC(w, r)
		return
	case http.MethodPost:
	default:
		shared.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

// Todo Get Use Case
func TodoGetUC(w http.ResponseWriter, r *http.Request) {
	page, limit, offset := shared.GetPagination(r)

	totalItems, err := shared.CountRows(
		db,
		"SELECT COUNT(*) FROM todos_tb",
	)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, "Failed count todo")
		return
	}

	rows, err := db.Query(
		`SELECT id, title, description, is_completed, due_date,
		        priority, created_at, updated_at, deleted_at
		 FROM todos_tb
		 ORDER BY id ASC
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, "Failed select todo")
		return
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Is_Completed,
			&todo.Due_Date,
			&todo.Priority,
			&todo.Created_At,
			&todo.Updated_At,
			&todo.Deleted_At,
		); err != nil {
			shared.RespondError(w, http.StatusInternalServerError, "Failed scan todo")
			return
		}
		todos = append(todos, todo)
	}

	meta := shared.Meta{
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: shared.CalculateTotalPages(totalItems, limit),
	}

	shared.RespondSuccess(w, http.StatusOK, "Get todo", todos, meta)
}
