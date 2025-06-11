package main

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *pgxpool.Pool
var ctx = context.Background()

func main() {
	log.Println("Starting application...")
	_ = godotenv.Load()
	connStr := os.Getenv("DATABASE_URL")
	var err error
	db, err = pgxpool.New(ctx, connStr)
	if err != nil {
		log.Printf("Database connection error: %v", err)
		panic(err)
	}
	log.Println("Successfully connected to database")
	defer db.Close()

	http.HandleFunc("/users", createUserHandler)
	http.HandleFunc("/users/list", listUsersHandler)
	http.HandleFunc("/users/delete", deleteUserHandler)
	http.HandleFunc("/users/update", updateUserHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Server listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Server startup error: %v", err)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request on /users with method: %s", r.Method)
	if r.Method != "POST" {
		log.Printf("Rejected request with unsupported method: %s", r.Method)
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	var input User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("JSON decoding error: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if input.Name == "" || input.Email == "" {
		log.Printf("Rejected request: name or email is empty (name: %s, email: %s)", input.Name, input.Email)
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	_, err := db.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", input.Name, input.Email)
	if err != nil {
		log.Printf("Database write error for user %s: %v", input.Name, err)
		http.Error(w, "Database write error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User added"))
	log.Printf("User with name %s and email %s was added!", input.Name, input.Email)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request on /users/list with method: %s", r.Method)
	if r.Method != "GET" {
		log.Printf("Rejected request with unsupported method: %s", r.Method)
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query(ctx, "SELECT name, email FROM users")
	if err != nil {
		log.Printf("Database query error: %v", err)
		http.Error(w, "Database read error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Name, &u.Email); err != nil {
			log.Printf("Row scan error: %v", err)
			http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	log.Printf("Successfully returned user list, total: %d", len(users))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request on /users/delete with method: %s", r.Method)
	if r.Method != "DELETE" {
		log.Printf("Rejected request with unsupported method: %s", r.Method)
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Email string `json:"Email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("JSON decoding error: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if input.Email == "" {
		log.Println("Rejected request: email string is empty")
		http.Error(w, "Email string is empty", http.StatusBadRequest)
		return
	}

	_, err := db.Exec(ctx, "DELETE FROM users WHERE email = $1", input.Email)
	if err != nil {
		log.Printf("Error deleting user with email %s: %v", input.Email, err)
		http.Error(w, "Error deleting user: ", http.StatusInternalServerError)
		return
	}
	log.Printf("User with email %s was successfully deleted!", input.Email)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request on /users/update with method: %s", r.Method)
	if r.Method != "PUT" {
		log.Printf("Rejected request with unsupported method: %s", r.Method)
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		OldEmail string `json:"OldEmail"`
		NewName  string `json:"name"`
		NewEmail string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("JSON decoding error: %v", err)
		http.Error(w, "JSON decoding error", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(ctx, "UPDATE users SET name = $1, email = $2 WHERE email = $3", input.NewName, input.NewEmail, input.OldEmail)
	if err != nil {
		log.Printf("Error updating user with email %s: %v", input.OldEmail, err)
		http.Error(w, "Error updating user data: ", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User data updated"))
	log.Printf("User with email %s was successfully updated! \n - New name: %s\n - New email: %s", input.OldEmail, input.NewName, input.NewEmail)
}
