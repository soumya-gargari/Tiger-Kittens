package main

import (
	"fmt"
	"log"
	"net/http"
	"tigerhall/tiger"
	"tigerhall/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// endpoint for creating user
	r.HandleFunc("/user", user.CreateUser).Methods("POST")
	// endpoint for user login
	r.HandleFunc("/login", user.Login).Methods("GET")
	// endpoint for creating tiger
	r.HandleFunc("/tiger", tiger.CreateTiger).Methods("POST")
	// endpoint for getting tiger details
	r.HandleFunc("/tiger", tiger.GetAllTigerDetails).Methods("GET")
	// endpoint for creating tiger sighting
	r.HandleFunc("/tiger/sighting", tiger.CreateTigerSighting).Methods("POST")
	// endpoint for getting tiger sighting details
	r.HandleFunc("/tiger", tiger.GetAllTigerSightingDetails).Methods("GET")
	fmt.Println("starting gorilla mux server:......")
	log.Fatal(http.ListenAndServe(":8080", r))
}
