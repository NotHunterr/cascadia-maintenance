package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


type Templates struct {
	templates *template.Template
}

func (t *Templates) RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("../views/*.html")),
	}
}

func validLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("Username: %s", username)
	fmt.Println("Password: ", password)
	fmt.Fprintf(w, "<p>Login Successful!!!!</p>")
}

func main() {
		router := chi.NewRouter()
		router.Use(middleware.Logger)
	
		// fs := http.FileServer(http.Dir("../views"))
		// router.Handle("/*", http.StripPrefix("/", fs))
		router.Get("/validLogin", validLoginHandler)
		fmt.Println("Server is running on port 42069")
		http.ListenAndServe("localhost:42069", router)
	}
