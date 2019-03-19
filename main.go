package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"main/src/permissions"
	"main/src/probe"
	"net/http"
	"os"
)

func _main(args []string) int {
	port := "80"
	if len(os.Getenv("PORT")) >= 2 {
		port = os.Getenv("PORT")
	}

	// Router
	router := mux.NewRouter().StrictSlash(true)

	// HealthCheck
	router.HandleFunc("/healthcheck", permissions.HealthCheck)

	// Probe
	router.HandleFunc("/probe", probe.Probe)

	// Tester
	router.HandleFunc("/tester", probe.Tester)

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
	fmt.Println(fmt.Sprintf("Starting Server on Port :%s", port))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		fmt.Println(fmt.Sprintf("HTTP: %v", err))
		return 1
	}

	fmt.Println("Died but nicely")
	return 0
}

func main() {
	os.Exit(_main(os.Args[1:]))
}