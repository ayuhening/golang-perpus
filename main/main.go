package main

import (
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/name/{name}", Greeting).Methods("GET")
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/book", setBook).Methods("POST")
	r.HandleFunc("/book", getBooks).Methods("GET")
	r.HandleFunc("/book/{ID}", delBook).Methods("DELETE")
	r.HandleFunc("/book/{ID}", updateBook).Methods("PUT")
	r.HandleFunc("/back", backRent).Methods("POST")
	r.HandleFunc("/rent", listBookforRent).Methods("GET")
	r.HandleFunc("/rent", rentBook).Methods("POST")
	r.HandleFunc("/user", userRegister).Methods("POST")
	r.HandleFunc("/user", getUser).Methods("GET")

	http.ListenAndServe(":8000", r)
}
