package main

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
	_ "net/http"
	"strconv"
	_ "strconv"
	"time"
	_ "time"
)

type Books struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	ISBN           string    `json:"isbn"`
	PublishingDate time.Time `json:"publishing_date"`
}

var books []Books

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Books
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = getNextID()
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func getNextID() int {
	if len(books) == 0 {
		return 1
	}
	return books[len(books)-1].ID + 1
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	for index, book := range books {
		if book.ID == id {
			var updatedBook Books
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = id
			books[index] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
