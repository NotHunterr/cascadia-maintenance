package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("../views"))
	router.Handle("/*", http.StripPrefix("/", fs))
	http.ListenAndServe(":42069", router)

}
