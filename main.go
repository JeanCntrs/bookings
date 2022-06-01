package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)

	http.ListenAndServe(portNumber, nil)
}
