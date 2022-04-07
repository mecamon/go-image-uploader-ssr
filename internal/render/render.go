package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"imageUploaderGo/internal/config"
	"imageUploaderGo/internal/models"
	"net/http"
	"path/filepath"
)

var appConfig *config.App
var globalPathToTemplates = "./templates"
var functions = template.FuncMap{}

func NewTemplate(a *config.App) {
	appConfig = a
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, templateData *models.TemplateData) error {

	var templateCached map[string]*template.Template
	if appConfig.UseCache {
		templateCached = appConfig.TemplateCache
	} else {
		templateCached, _ = CreateTemplateCached()
	}

	t, ok := templateCached[tmpl]
	if !ok {
		return errors.New("error receiving template from appConfig")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, templateData)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error printing template to browser:", err)
		return err
	}

	return nil
}

func CreateTemplateCached() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", globalPathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", globalPathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			_, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", globalPathToTemplates))
			if err != nil {
				return myCache, nil
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
