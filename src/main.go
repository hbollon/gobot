package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Declare mux router and out port
	router := mux.NewRouter()
	port := getPort()

	// Declare the HTTP methods for our routes
	router.HandleFunc("/", indexHandler).Methods("GET")
	err := http.ListenAndServe(port, router)

	// Error handling on server startup
	if err != nil {
		log.Fatal("Error on server startup: ", err)
	}
}

// Manage index route (/)
func indexHandler(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprint(rw, "Server is up and functional !")
}

// Return port
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3500"
		fmt.Printf("WARNING : Port not defined, using default port (%s)\n", port)
	}
	return port
}
