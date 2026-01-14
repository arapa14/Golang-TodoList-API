package main

import (
	"TODO-LIST-API/infrastructure/database"
	"TODO-LIST-API/internal/config"
	"TODO-LIST-API/internal/shared"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// Data table 1: Todo
type Todo struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  *string `json:"description"`
	Is_Completed bool   `json:"is_completed"`
	Due_Date     *string `json:"due_date"`
	Priority     int    `json:"priority"`
	Created_At   string `json:"created_at"`
	Updated_At   string `json:"updated_at"`
	Deleted_At   *string `json:"deleted_at"`
}

// DTO
type PostTodo struct {
	Title        string `json:"title"`
	Description  *string `json:"description"`
	Due_Date     *string `json:"due_date"`
	Priority     int    `json:"priority"`
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
		TodoPostUC(w, r)
		return
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

// Todo Post Use Case
func TodoPostUC(w http.ResponseWriter, r *http.Request) {
	var postTodo PostTodo
	var newTodo Todo

	if err := json.NewDecoder(r.Body).Decode(&postTodo); err != nil {
		shared.RespondError(w, http.StatusBadRequest, "JSON isn't valid")
		return
	}

	if postTodo.Title == "" {
		shared.RespondError(w, http.StatusBadRequest, "Title cannot empty")
		return
	}

	err := db.QueryRow(
		"INSERT INTO todos_tb (title, description, priority, due_date) VALUES ($1, $2, $3, $4) RETURNING id, title, description, is_completed, due_date, priority, created_at, updated_at, deleted_at",
		postTodo.Title, postTodo.Description, postTodo.Priority, postTodo.Due_Date,
	).Scan(&newTodo.ID, &newTodo.Title, &newTodo.Description, &newTodo.Is_Completed, &newTodo.Due_Date, &newTodo.Priority, &newTodo.Created_At, &newTodo.Updated_At, &newTodo.Deleted_At)
	if err != nil {
		log.Println(err)
		shared.RespondError(w, http.StatusInternalServerError, "Failed insert todo")
		return
	}

	shared.RespondSuccess(w, http.StatusCreated, "Post todo", newTodo)
	return
}