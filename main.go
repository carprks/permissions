package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"main/src/permissions"
	"net/http"
	"os"
)

func main() {
	// Router
	router := mux.NewRouter().StrictSlash(true)

	// HealthCheck
	router.HandleFunc("/healthcheck", permissions.HealthCheck)

	// General Permissions
	router.HandleFunc("/", permissions.CreateRoute).Methods("POST")
	router.HandleFunc("/", permissions.RetrieveRoute).Methods("GET")
	router.HandleFunc("/", permissions.UpdateRoute).Methods("PATCH")
	router.HandleFunc("/", permissions.DeleteRoute).Methods("DELETE")

	// User Permissions
	router.HandleFunc("/user", permissions.CreateUserRoute).Methods("POST")
	router.HandleFunc("/user", permissions.RetrieveUserRoute).Methods("GET")
	router.HandleFunc("/user", permissions.UpdateUserRoute).Methods("PATCH")
	router.HandleFunc("/user", permissions.DeleteUserRoute).Methods("DELETE")

	// Start Server
	if os.Getenv("PORT") != "" {
		fmt.Println(fmt.Sprintf("Starting Server on Port :%s", os.Getenv("PORT")))
		log.Fatal("Start Server", http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
	} else {
		log.Fatal("No Port Specified")
	}
}
