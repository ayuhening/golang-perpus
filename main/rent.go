package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//fixed
func listBookforRent(w http.ResponseWriter, r *http.Request) {
	var book_list []Books
	var books Books
	var response Response

	db := database()
	defer db.Close()

	rows, err := db.Query("Select * from book where IsReady = true")

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
func backRent(w http.ResponseWriter, r *http.Request) {
	var books Books
	var rents Rent
	var users User
	var response Response
	var daterent = time.Now()

	db := database()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&rents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Query("SELECT * FROM book WHERE ID = ?", rents.BookID)
	if err != nil {
		log.Print(err)
	}

	for result.Next() {
		err := result.Scan(&books.ID, &books.Title, &books.Author, &books.IsReady)
		if err != nil {
			log.Print(err)
		}
	}

	if rents.BookID != books.ID || rents.UserID != users.ID {
		response.Status = 1
		response.Message = "Book ID / User ID is invalid, please check again!"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)

		//Tanyakan ke Ferdi, kenapa kalau masukin UserID udah bener tp BookID yg salah, tetep munculin error ini
	} else if books.IsReady == false {

		result, err := db.Query("SELECT * FROM rent WHERE BookID = ?", rents.BookID)
		if err != nil {
			log.Print(err)
		}

		for result.Next() {
			err := result.Scan(&rents.ID, &rents.UserID, &rents.BookID, &rents.RentDate, &rents.BackDate, &rents.RealTimeBack)
			if err != nil {
				log.Print(err)
			}
		}

		tx, err := db.Begin()
		if err != nil {
			return
		}

		result, err = db.Query("UPDATE rent set RealTimeBack = ? WHERE BookID = ?", daterent, rents.BookID)
		result, err = db.Query("UPDATE book set IsReady = true WHERE ID = ?", rents.BookID)

		if err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()

		rents.RealTimeBack = fmt.Sprintf("%s", daterent.Format(time.RFC3339))

		response.Status = 1
		response.Message = "Book returned!"
		response.Data = &rents
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	} else {
		response.Status = 1
		response.Message = "Book is not borrowed, check the ID again!"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
	}
}

//fixed
func rentBook(w http.ResponseWriter, r *http.Request) {
	var books Books
	var user User
	var response Response
	var rents Rent
	var daterent = time.Now()
	var dateback = daterent.Add(time.Hour * 24 * 7)

	db := database()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&rents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Query("SELECT * FROM book WHERE ID = ?", rents.BookID)
	if err != nil {
		log.Print(err)
	}

	for result.Next() {
		err := result.Scan(&books.ID, &books.Title, &books.Author, &books.IsReady)
		if err != nil {
			log.Print(err)
		}
	}

	if books.ID != rents.BookID {
		response.Status = 1
		response.Message = "Book with ID " + rents.BookID + " is not available."

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)

	} else if books.IsReady == true {
		tx, err := db.Begin()
		if err != nil {
			return
		}

		_, err = tx.Query("UPDATE book set IsReady = false WHERE ID = ?", rents.BookID)

		result, err := tx.Exec("INSERT INTO rent (UserID, BookID, RentDate, BackDate) values (?, ?, ?, ?)",
			rents.UserID,
			rents.BookID,
			daterent,
			dateback,
		)
		if err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rents.ID = fmt.Sprintf("%d", id)
		rents.RentDate = fmt.Sprintf("%s", daterent.Format(time.RFC3339))
		rents.BackDate = fmt.Sprintf("%s", dateback.Format(time.RFC3339))

		response.Status = 1
		response.Message = "Book borrowed!"
		response.Data = &rents

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	} else {
		response.Status = 1
		response.Message = "Book with ID " + rents.BookID + " was borrowed."
		response.Data = &books

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
	}

	if user.ID != rents.UserID {
		response.Status = 1
		response.Message = "User with ID " + rents.UserID + " is not available."

	}
}
