package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("../views"))
	router.Handle("/*", http.StripPrefix("/", fs))
	router.Get("/validLogin", validLoginHandler)
	http.ListenAndServe("localhost:42069", router)

}

func validLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "admin" {
		fmt.Println("Login successful")
	}
}
