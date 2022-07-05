package main

type Books struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Author  string `json:"Author"`
	IsReady bool   `json:"IsReady"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"Data"`
}

type User struct {
	ID           string `json:"ID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	EmailAddress string `json:"EmailAddress"`
}

type Rent struct {
	ID           string `json:"ID"`
	UserID       string `json:"UserID"`
	BookID       string `json:"BookID"`
	RentDate     string `json:"RentDate"`
	BackDate     string `json:"BackDate"`
	RealTimeBack string `json:"RealTimeBack"`
}
