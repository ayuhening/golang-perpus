package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func userRegister(w http.ResponseWriter, r *http.Request) {
	var users User

	db := database()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM user WHERE EmailAddress = ?", users.EmailAddress)

	err = row.Scan(&users.ID, &users.FirstName, &users.LastName, &users.EmailAddress)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err == nil {
		json.NewEncoder(w).Encode("email registered.")
		return
	}

	rows, err := db.Exec("INSERT INTO user (FirstName, LastName, EmailAddress) values (?, ?, ?)",
		users.FirstName,
		users.LastName,
		users.EmailAddress,
	)
	if err != nil {
		log.Print(err)
		return
	}

	id, err := rows.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users.ID = fmt.Sprintf("%d", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("user added.")

}

func getUser(w http.ResponseWriter, r *http.Request) {
	var user_list []User
	var users User
	var response Response

	db := database()
	defer db.Close()

	rows, err := db.Query("Select * from user")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.FirstName, &users.LastName, &users.EmailAddress); err != nil {
			log.Fatal(err.Error())
		} else {
			user_list = append(user_list, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = user_list

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
