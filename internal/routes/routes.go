package routes

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"imageUploaderGo/internal/models"
	"imageUploaderGo/internal/render"
	"log"
	"net/http"
)

func MakeRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/", home)
	router.Post("/upload", upload)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Fatal(err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Url string `json:"url"`
	}{Url: "some-random-url"}

	out, _ := json.MarshalIndent(resp, "", "   ")
	w.Write(out)
}
