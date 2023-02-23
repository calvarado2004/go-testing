package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

var pathToTemplates = "./templates/"

// Home is the handler for the home page
func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	err := app.render(w, r, "home.page.gohtml", &TemplateData{})
	if err != nil {
		log.Printf("error rendering template: %v", err)
	}

}

type TemplateData struct {
	IP   string
	Data map[string]any
}

// render is a helper function that parses a template file and writes the
func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {

	// parse the template from disk
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	// write the template to the http.ResponseWriter
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return err
	}

	return nil
}

// Login is the handler for the login page
func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println("error parsing form: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")

	password := r.Form.Get("password")

	log.Printf("email: %v, password %v", email, password)

	fmt.Fprintf(w, "email: %v", email)

}
