package user

import (
	"Tiger-Kittens/data"
	"Tiger-Kittens/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CreateUser method for creating user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	var userDetails data.UserDetails
	err = json.Unmarshal(body, &userDetails)
	if err != nil {
		http.Error(w, "failed to unmarshal", http.StatusBadRequest)
		return
	}
	var dB database.Database
	err = dB.CreateUserTable(data.UserTableName)
	if err != nil {
		http.Error(w, "failed to save data to mysql", http.StatusInternalServerError)
		return
	}
	err = dB.InsertUserData(data.UserTableName, userDetails)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to save data to mysql Table", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(http.StatusCreated)
}

// Login method for user Login
func Login(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	queryParams := r.URL.Query()

	// Get the value of the "username" query parameter
	username := queryParams.Get("username")

	// Check if the parameter is present
	if username == "" {
		http.Error(w, "Missing 'username' query parameter", http.StatusBadRequest)
		return
	}

	// Get the value of the "password" query parameter
	password := queryParams.Get("password")

	// Check if the parameter is present
	if password == "" {
		http.Error(w, "Missing 'password' query parameter", http.StatusBadRequest)
		return
	}
	db := &database.Database{}
	mySql := db.InitDatabse()
	db.Db = mySql
	userDetails, err := db.GetUserData(data.UserTableName, username, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to login", http.StatusInternalServerError)
		return
	}
	if userDetails.UserName != username || userDetails.PassWord != password {
		http.Error(w, "failed to login", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(http.StatusOK)
}
