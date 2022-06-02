package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JeanCntrs/bookings/pkg/config"
	"github.com/JeanCntrs/bookings/pkg/handlers"
	"github.com/JeanCntrs/bookings/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}
