package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akirayorunoe/go-course/pkg/config"
	"github.com/akirayorunoe/go-course/pkg/handlers"
	"github.com/akirayorunoe/go-course/pkg/render"
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

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
