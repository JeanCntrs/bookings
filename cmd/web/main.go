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
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
