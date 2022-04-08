package main

import (
	"imageUploaderGo/internal/config"
	"imageUploaderGo/internal/render"
	"imageUploaderGo/internal/routes"
	"log"
	"net/http"
)

var portNumber = ":8080"
var appConfig config.App

func main() {

	templateCached, err := render.CreateTemplateCached()
	if err != nil {
		log.Fatal(err)
	}
	appConfig.TemplateCache = templateCached
	appConfig.UseCache = true

	render.NewTemplate(&appConfig)

	handler := routes.MakeRouter()

	log.Printf("Server running on port %s...", portNumber)

	log.Fatal(http.ListenAndServe(portNumber, handler))
}

func GetAppConfig() config.App {
	return appConfig
}
