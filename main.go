package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book represents the data model for a book
type Book struct {
	ID     int    `json:"id"`     // Unique identifier for the book
	Title  string `json:"title"`  // Title of the book
	Author string `json:"author"` // Author of the book
	Year   int    `json:"year"`   // Year of publication
}

// books is an in-memory data store simulating a database
var books []Book

// autoIncrementID keeps track of the next ID to assign to a new book
var autoIncrementID = 1

// getBooks handles GET /books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// getBook handles GET /books/{id}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

// createBook handles POST /books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	newBook.ID = autoIncrementID
	autoIncrementID++
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// updateBook handles PUT /books/{id}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedBook Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.NotFound(w, r)
}

// deleteBook handles DELETE /books/{id}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

// seedData initializes the books slice with sample data
func seedData() {
	books = []Book{
		{ID: autoIncrementID, Title: "1984", Author: "George Orwell", Year: 1949},
	}
	autoIncrementID++
	books = append(books, Book{ID: autoIncrementID, Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960})
	autoIncrementID++
	books = append(books, Book{ID: autoIncrementID, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925})
	autoIncrementID++
}

// handleRoutes defines all API routes and handlers
func handleRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// main is the entry point of the application
func main() {
	seedData()
	handleRoutes()
}
