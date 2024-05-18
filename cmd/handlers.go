package main

import (
	"html/template"
	"net/http"
	"path"
)

type TemplateData struct {
	Data map[string]any
}

var pathToTemplates = "./templates/"

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData, isLogged bool) error {

	if !isLogged {
		parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "/bases/base.layout.gohtml"))
		if err != nil {
			return err
		}
		err = parsedTemplate.Execute(w, td)
		if err != nil {
			return err
		}
		return nil
	}

	// if isLogged true
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "/bases/logged.layout.gohtml"))
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
	_ = app.render(w, r, "acceuil.gohtml", &TemplateData{}, false)
}
