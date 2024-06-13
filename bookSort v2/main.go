package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"
	"log"
	_ "log"
	"sort"
	"time"
	_ "time"

	_ "modernc.org/sqlite"
)

type Book struct {
	ID     int64
	Title  string
	Author string
	ISBN   int
	Date   time.Time
}

func main() {
	db, err := sql.Open("sqlite", "books.db")
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		isbn INTEGER,
		date DATE
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	books := []Book{
		{Title: "How to teach for dummies", Author: "Alan West", ISBN: 978012653240, Date: time.Now()},
		{Title: "Sherlock", Author: "Arthur C. Doyle", ISBN: 978032516224, Date: time.Now()},
		{Title: "Clean house", Author: "Roberta Gonzalez", ISBN: 9780132350884, Date: time.Now()},
	}

	for _, book := range books {
		insertBook(db, book)
	}

	retrievedBooks, err := retrieveBooks(db)
	if err != nil {
		log.Fatalf("failed to retrieve books: %v", err)
	}

	sort.SliceStable(retrievedBooks, func(i, j int) bool {
		return retrievedBooks[i].Title < retrievedBooks[j].Title
	})

	fmt.Println("Sorted Books by Title:")
	for _, book := range retrievedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s, ISBN: %d, Date: %s\n", book.ID, book.Title, book.Author, book.ISBN, book.Date.Format("2006-01-02"))
	}
}
func insertBook(db *sql.DB, book Book) {
	insertQuery := `INSERT INTO books (title, author, isbn, date) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(insertQuery, book.Title, book.Author, book.ISBN, book.Date.Format(time.RFC3339))
	if err != nil {
		log.Fatalf("failed to insert book: %v", err)
	}
}
func retrieveBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author, isbn, date FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		var dateStr string
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &dateStr); err != nil {
			return nil, err
		}
		book.Date, _ = time.Parse(time.RFC3339, dateStr)
		books = append(books, book)
	}
	return books, nil
}
