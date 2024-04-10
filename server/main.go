package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// To-do - learn enum for locations, might make it easier

type CreateNewTicket struct {
	title string
	description string
	location string
}

func validLoginHandler(w http.ResponseWriter, r *http.Request) {

		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Printf("Username: %s", username)
		fmt.Println("Password: ", password)

		if username == "admin" && password == "admin" {
			fmt.Fprintf(w, "Login Successful! Re-directing")
			fmt.Fprintf(w, "<script>setTimeout(function(){window.location.href = '../locations/tickets/ViewTickets.html';}, 1000);</script>")
	} else {
		fmt.Fprintf(w, "Login Failed! Please try again.")
		fmt.Fprintf(w, "<script>setTimeout(function(){window.location.href = './login.html';}, 1500);</script>")

	}

}

func createNewTicketHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("create-title")
	description := r.FormValue("create-description")
	location := r.FormValue("create-location")

	fmt.Printf("title: %s, description: %s, location: %s", title, description, location)



}


	func main() {
		// Initializing Chi router
		router := chi.NewRouter()
		router.Use(middleware.Logger)
		
		// Initializing ENV variables
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// Initializing mysql database
		sqlpw := os.Getenv("MYSQLPW")
		db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/cascadia_maintenance", sqlpw))
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		fmt.Println("Success")

	
		fs := http.FileServer(http.Dir("../views"))
		router.Handle("/*", http.StripPrefix("/", fs))
		router.Get("/validLogin", validLoginHandler)
		router.Post("/createNewTicket", createNewTicketHandler)
		fmt.Println("Server is running on port 42069")
		http.ListenAndServe("localhost:42069", router)
	}
