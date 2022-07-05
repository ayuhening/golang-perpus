package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/gorilla/mux"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type Book struct {
// 	ID      string `json:"ID"`
// 	Title   string `json:"Title"`
// 	Author  string `json:"Author"`
// 	IsReady bool   `json:"IsReady"`
// }

// var db *sql.DB
// var err error

// func Greeting(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	name := params["name"]
// 	w.Write([]byte("Hi "))
// 	w.Write([]byte(name))
// 	w.Header().Set("Content-Type", "application/json")
// }

// func Home(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Welcome!"))
// 	w.Header().Set("Content-Type", "application/json")
// }

// func main() {
// 	db, err = sql.Open("mysql", "root:Ayuhening1997!((&@tcp(127.0.0.1:3306)/Perpus")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	r := mux.NewRouter()
// 	r.HandleFunc("/name/{name}", Greeting).Methods("GET")
// 	r.HandleFunc("/", Home).Methods("GET")
// 	r.HandleFunc("/book", setBook).Methods("POST")
// 	r.HandleFunc("/book", getBooks).Methods("GET")
// 	r.HandleFunc("/book/{ID}", getBook).Methods("GET")
// 	//r.HandleFunc("/book/{ID}", delBook).Methods("DELETE")
// 	//r.HandleFunc("/book/{ID}", updateBook).Methods("PUT")
// 	http.ListenAndServe(":8000", r)

// }

// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var books []Book

// 	result, err := db.Query("SELECT ID, Title, Author, IsReady from book")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()

// 	for result.Next() {
// 		var book Book
// 		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.IsReady)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		books = append(books, book)
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	result, err := db.Query("SELECT ID, Title, Author, IsReady FROM book WHERE ID = ?", params["ID"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	var book Book
// 	for result.Next() {
// 		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.IsReady)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}

// 	json.NewEncoder(w).Encode(book)

// }
