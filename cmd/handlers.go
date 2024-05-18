package main

import (
	"html/template"
	"net/http"
	"path"
)

type TemplateData struct {
	Data map[string]any
}

var pathToTemplates = "./static/templates/"

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {

	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "base.layout.gohtml"))
	if err != nil {
		return err
	}
	err = parsedTemplate.Execute(w, td)
	if err != nil {
		return err
	}
	return nil
}

// template rendering

func (app *application) RenderAccueil(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "acceuil.gohtml", &TemplateData{})
}

func (app *application) RenderLoged(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "loged.gohtml", &TemplateData{})
}
