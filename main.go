package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"main/src/healthcheck"
	"main/src/probe"
	"net/http"
	"os"
)

func _main(args []string) int {
	// development
	if len(args) >= 1 {
		if args[0] == "localDev" {
			godotenv.Load()
		}
	}

	port := "80"
	if len(os.Getenv("PORT")) >= 2 {
		port = os.Getenv("PORT")
	}

	// Router
	router := mux.NewRouter().StrictSlash(true)

	// healthcheck
	router.HandleFunc("/healthcheck", healthcheck.HTTP)

	// Probe
	router.HandleFunc("/probe", probe.HTTP)

	// User Permissions
	// router.HandleFunc("/users/", permissions.CreateUser).Methods("POST")
	// router.HandleFunc("/users/", permissions.RetrieveAllUsers).Methods("GET")
	// router.HandleFunc("/users/{permission}/", permissions.RetrieveUser).Methods("GET")
	// router.HandleFunc("/users/{permission}/", permissions.UpdateUser).Methods("PATCH")
	// router.HandleFunc("/users/{permission}/", permissions.DeleteUser).Methods("DELETE")

	// General Permission
	// router.HandleFunc("/", permissions.Create).Methods("POST")
	// router.HandleFunc("/", permissions.RetrieveAll).Methods("GET")
	// router.HandleFunc("/{permission}/", permissions.RetrievePermissions).Methods("GET")
	// router.HandleFunc("/{permission}/", permissions.UpdatePermission).Methods("PATCH")
	// router.HandleFunc("/{permission}/", permissions.DeletePermission).Methods("DELETE")

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