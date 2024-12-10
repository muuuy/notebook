/*
	Handles routes for Users
*/

package handlers

import (
	"encoding/json"
	"net/http"

	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// func enableCors(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins (use specific origins for production)
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allowed HTTP methods
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allowed headers

// 	if r.Method == http.MethodOptions {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}
// }

// Set up ID and User store
var nextID = 1
var users = []User{}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	w.Write([]byte("User ID: " + idString))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")              // Allow all origins (use specific origins for production)
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // Allowed HTTP methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")  // Allowed headers

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("Why isn't this shit working?")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if newUser.Name == "" || newUser.Email == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}

	// Assign new ID & Increment ID for next user
	newUser.ID = nextID
	nextID++

	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func RegisterUserRoutes() {
	http.HandleFunc("GET /users", GetUsers)
	http.HandleFunc("GET /users/{id}", GetUserByID)
	http.HandleFunc("POST /users", CreateUser)
}
