package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Information struct {
	Information string
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("../views/authentication/*.html")),
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
		router := echo.New()
		router.Use(middleware.Logger())
	
		// fs := http.FileServer(http.Dir("../views"))
		// router.Handle("/*", http.StripPrefix("/", fs))

		router.Renderer = newTemplates()
		// fmt.Println("Server is running on port 42069")
		// http.ListenAndServe("localhost:42069", router)
		router.GET("/", func(c echo.Context) error {
			return c.Render(200, "../views/authentication/login.html", nil)
		})

	}
