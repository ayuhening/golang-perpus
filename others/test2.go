package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type book struct {
// 	ID      string `json:"ID"`
// 	Title   string `json:"Title"`
// 	Author  string `json:"Author"`
// 	IsReady bool   `json:"IsReady"`
// }

// `json:"-"` >> buat meng-ignore || `json:"Author,omitempty"` >> untuk tidak ngeprint kalau value-nya kosong

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

// var book_list []book

// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(book_list)
// }

// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	var status int

// 	for _, item := range book_list {

// 		if item.ID == params["ID"] {
// 			status, _ = strconv.Atoi(item.ID)
// 			json.NewEncoder(w).Encode(item)
// 			break
// 		} else if item.ID != params["ID"] {
// 			status = 0
// 		}
// 	}

// 	if status == 0 {
// 		json.NewEncoder(w).Encode("The book doesn't exist.")
// 	}
// }

// func setBook(w http.ResponseWriter, r *http.Request) {
// 	var book book
// 	json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(len(book_list) + 1)
// 	book_list = append(book_list, book)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("Book is successfully added!")
// }

// func delBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["ID"]
// 	for index, book := range book_list {
// 		if book.ID == id {
// 			book_list = append(book_list[:index], book_list[index+1:]...)
// 		}
// 	}
// }

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range book_list {
// 		if item.ID == params["ID"] {
// 			book_list = append(book_list[:index], book_list[index+1:]...)
// 			var book book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["ID"]
// 			book_list = append(book_list, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}

// }

// func handleRequest() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/name/{name}", Greeting).Methods("GET")
// 	r.HandleFunc("/", Home).Methods("GET")
// 	r.HandleFunc("/book", setBook).Methods("POST")
// 	r.HandleFunc("/book", getBooks).Methods("GET")
// 	r.HandleFunc("/book/{ID}", getBook).Methods("GET")
// 	r.HandleFunc("/book/{ID}", delBook).Methods("DELETE")
// 	r.HandleFunc("/book/{ID}", updateBook).Methods("PUT")
// 	log.Fatal(http.ListenAndServe(":7000", r))

// }

// func main() {
// 	book_list = []book{}

// 	handleRequest()
// }
