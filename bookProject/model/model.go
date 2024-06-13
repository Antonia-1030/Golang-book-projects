package model

import (
	"bookProject/db"
	_ "bookProject/db"
	"time"
)

type Book struct {
	Id     int64
	Title  string
	Author string
	ISBN   int
	Date   time.Time
}

var bookCollection []Book

func (b *Book) Save() error {

	statment, err := db.GetDb().Prepare(`
		INSERT INTO 
		movies
		    (title, author, isbn,date)
		VALUES
		    (?, ?, ?,?)
	`)

	defer statment.Close()

	if err != nil {
		return err
	}

	result, err := statment.Exec(b.Title, b.Author, b.ISBN, b.Date)
	if err != nil {
		return err
	}

	movieId, err := result.LastInsertId()
	b.Id = movieId

	return err
}

func GetAllBooks() ([]Book, error) {

	dbCursor, err := db.GetDb().Query(`SELECT * FROM books`)
	if err != nil {
		return nil, err
	}

	for dbCursor.Next() {

		var movieObject Book
		err := dbCursor.Scan(
			&movieObject.Id,
			&movieObject.Title,
			&movieObject.Author,
			&movieObject.ISBN,
			&movieObject.Date,
		)

		if err != nil {
			return nil, err
		}

		bookCollection = append(bookCollection, movieObject)
	}

	return bookCollection, nil
}
