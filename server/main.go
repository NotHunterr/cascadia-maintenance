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

func validLoginHandler(w http.ResponseWriter, r *http.Request) {

		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Printf("Username: %s", username)
		fmt.Println("Password: ", password)

		if username == "admin" && password == "admin" {
			fmt.Fprintf(w, "Login Successful! Re-directing")
			fmt.Fprintf(w, "<script>setTimeout(function(){window.location.href = '../locations/AssemblyFloor.html';}, 1000);</script>")
	} else {
		fmt.Fprintf(w, "Login Failed! Please try again.")
		fmt.Fprintf(w, "<script>setTimeout(function(){window.location.href = './login.html';}, 1500);</script>")

	}

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
		fmt.Println("Server is running on port 42069")
		http.ListenAndServe("localhost:42069", router)
	}
