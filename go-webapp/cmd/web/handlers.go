package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var pathToTemplates = "./templates/"

// Home is the handler for the home page
func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	var td = make(map[string]any)

	if app.Session.Exists(r.Context(), "test") {
		message := app.Session.GetString(r.Context(), "test")
		td["test"] = message
		log.Printf("session exists, message: %v", message)
	} else {
		app.Session.Put(r.Context(), "test", "Hit this page at "+time.Now().UTC().String())
		log.Printf("session created, it was empty")
	}

	err := app.render(w, r, "home.page.gohtml", &TemplateData{
		Data: td,
	})
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
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "base.layout.gohtml"))
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
		log.Printf("error parsing form: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// validate the form data
	form := NewForm(r.PostForm)
	form.Required("email", "password")

	if !form.Valid() {
		log.Printf("form not valid: %v", form.Errors)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.DB.GetUserByEmail(email)
	if err != nil {
		log.Printf("error getting user by email: %v", err)
		return
	}

	log.Println("From DB: ", user)

	form.IsEmail(email)

	form.MinLength(password, 3)

	log.Printf("email: %v, password %v", email, password)

	fmt.Fprintf(w, "email: %v", email)

}
