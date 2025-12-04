package main

import (
	// "fmt"
	"fmt"

	// "log"
	"net/http"

	"example.com/http-server/internal/handlers"
	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {

	router.HandleFunc("/hello", handlers.GetHandler).Methods("GET")
	router.HandleFunc("/echo", handlers.PostHandler).Methods("POST")
	router.HandleFunc("/hello/{name}", handlers.PathParams).Methods("GET")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

}

func main() {
	r := mux.NewRouter()
	Routes(r)
	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", r)	
}
