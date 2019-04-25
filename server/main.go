package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// book struct == Model
type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

// author struct
type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname`
}

//login struct
type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

// Init books var as a slice
var books []Book
var users []User

// err checker
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) // get the params
	// loops through books and find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100)) // Mock Id not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	// w.Header().Set("Content-type", "application/json")
	// var user User
	// _ = json.NewDecoder(r.Body).Decode(&user)
	// checkErr(err)
	// res, err := stmt.Exec(user.Username, user.Password)
	// checkErr(err)
	// json.NewEncoder(w).Encode(user)
}

func main() {
	// init new router
	r := mux.NewRouter()

	// Mock Data # implement database
	books = append(books, Book{ID: "1", Isbn: "5452523523", Title: "The garden book", Author: &Author{Firstname: "Richason", Lastname: "Tjongirin "}})
	books = append(books, Book{ID: "2", Isbn: "312312312123", Title: "The dota book", Author: &Author{Firstname: "woof", Lastname: "boi "}})
	books = append(books, Book{ID: "3", Isbn: "8678678", Title: "Game of thrones", Author: &Author{Firstname: "jeka", Lastname: "rowling "}})

	// route handlers
	// r.HandleFunc("/api/login", loginUser)
	r.HandleFunc("/api/register", registerUser).Methods("POST")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
