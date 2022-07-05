package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//fixed
func getBooks(w http.ResponseWriter, r *http.Request) {
	var book_list []Books
	var books Books
	var response Response

	db := database()
	defer db.Close()

	rows, err := db.Query("Select * from book")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&books.ID, &books.Title, &books.Author, &books.IsReady); err != nil {
			log.Fatal(err.Error())
		} else {
			book_list = append(book_list, books)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = book_list

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

//fixed
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var books Books
	var response Response

	db := database()
	defer db.Close()

	result, err := db.Query("SELECT * FROM book WHERE ID = ?", params["ID"])
	if err != nil {
		log.Print(err)
	}

	err = result.Scan(&books.ID, &books.Title, &books.Author, &books.IsReady)
	if err != nil {
		log.Print(err)
	}

	if params["ID"] != books.ID {
		response.Status = 1
		response.Message = "Book with ID " + params["ID"] + " is not available."
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
	} else {
		response.Status = 1
		response.Message = "Success"
		response.Data = &books
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

//fixed
func setBook(w http.ResponseWriter, r *http.Request) {
	var books Books
	var response Response
	//var datetime = time.Now()

	db := database()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO book (Title, Author, IsReady) values (?, ?, ?)",
		books.Title,
		books.Author,
		books.IsReady,
		//datetime,
	)
	if err != nil {
		log.Print(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books.ID = fmt.Sprintf("%d", id)
	//books.RentDate = fmt.Sprintf("%s", datetime)

	response.Status = 1
	response.Message = "Book added!!!"
	response.Data = &books

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//fixed
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var books Books
	var response Response

	db := database()
	defer db.Close()

	result, err := db.Query("SELECT ID, Title, Author, IsReady FROM book WHERE ID = ?", params["ID"])
	if err != nil {
		log.Print(err)
	}

	for result.Next() {
		err := result.Scan(&books.ID, &books.Title, &books.Author, &books.IsReady)
		if err != nil {
			log.Print(err)
		}
	}

	if params["ID"] != books.ID {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("Book with ID " + params["ID"] + " is not available.")
	} else {
		err := json.NewDecoder(r.Body).Decode(&books)

		if err != nil {
			log.Print(err)
		}

		_, err = db.Exec("UPDATE book set Title = ?, Author = ?, IsReady = ? WHERE ID = ?",
			books.Title,
			books.Author,
			books.IsReady,
			params["ID"],
		)

		if err != nil {
			log.Print(err)
		}

		response.Status = 1
		response.Message = "Success"
		response.Data = &books
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

//fixed
func delBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var books Books
	var response Response

	db := database()
	defer db.Close()

	result, err := db.Query("SELECT ID, Title, Author, IsReady FROM book WHERE ID = ?", params["ID"])
	if err != nil {
		log.Print(err)
	}

	for result.Next() {
		err := result.Scan(&books.ID, &books.Title, &books.Author, &books.IsReady)
		if err != nil {
			log.Print(err)
		}
	}

	if params["ID"] != books.ID {
		response.Status = 1
		response.Message = "Book with ID " + params["ID"] + " is not available."

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
	} else {
		_, err = db.Exec("DELETE from book where ID = ?", params["ID"])

		response.Status = 1
		response.Message = "Successfully removed."
		response.Data = &books

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
