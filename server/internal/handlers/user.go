/*
	Handles routes for Users
*/

package handlers

import (
	"net/http"
)

func RegisterUserRoutes() {
	http.HandleFunc("GET /users", GetUsers)
	http.HandleFunc("GET /users/{id}", GetUserByID)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	w.Write([]byte("User ID: " + idString))
}
