package config

import "html/template"

type App struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
}
