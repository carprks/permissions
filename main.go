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

	// User Permissions
	router.HandleFunc("/users/", permissions.CreateUserRoute).Methods("POST")
	router.HandleFunc("/users/", permissions.RetrieveUserAllRoute).Methods("GET")
	router.HandleFunc("/users/{permission}/", permissions.RetrieveUserRoute).Methods("GET")
	router.HandleFunc("/users/{permission}/", permissions.UpdateUserRoute).Methods("PATCH")
	router.HandleFunc("/users/{permission}/", permissions.DeleteUserRoute).Methods("DELETE")

	// General Permission
	router.HandleFunc("/", permissions.CreateRoute).Methods("POST")
	router.HandleFunc("/", permissions.RetrieveAllRoute).Methods("GET")
	router.HandleFunc("/{permission}/", permissions.RetrieveRoute).Methods("GET")
	router.HandleFunc("/{permission}/", permissions.UpdateRoute).Methods("PATCH")
	router.HandleFunc("/{permission}/", permissions.DeleteRoute).Methods("DELETE")

	// Start Server
	if os.Getenv("PORT") != "" {
		fmt.Println(fmt.Sprintf("Starting Server on Port :%s", os.Getenv("PORT")))
		log.Fatal("Start Server", http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
	} else {
		log.Fatal("No Port Specified")
	}
}
