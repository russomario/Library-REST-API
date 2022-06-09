//    Copyright 2022 Russo Mario mario@russomario.xyz

//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at

//        http://www.apache.org/licenses/LICENSE-2.0

//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package database_manager

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Book struct {
	ISBN    string `json:"isbn" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Summary string `json:"summary" binding:"required"`
	PubYear int    `json:"pub_year" binding:"required"`
}

type Author struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Birth string `json:"birth"`
}

type dbHandler struct {
	handler         *sql.DB
	name            string
	databaseName    string
	databaseAddress string
	databasePort    int
}

var db = dbHandler{
	name:            "LibraryHandler",
	databaseName:    "library",
	databaseAddress: "localhost",
	databasePort:    3306,
}

func SetUpDBConnection() {

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%v:%v", db.databaseAddress, db.databasePort),
		DBName: db.databaseName,
	}

	// Get a database handle.
	var err error
	db.handler, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.handler.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func GetBooks() ([]Book, error) {
	var books []Book

	rows, err := db.handler.Query("SELECT * FROM book")
	if err != nil {
		return nil, fmt.Errorf("GetBooks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ISBN, &book.Title, &book.Author, &book.Summary, &book.PubYear); err != nil {
			return nil, fmt.Errorf("GetBooks: %v", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBooks: %v", err)
	}
	return books, nil
}

func GetBookByISBN(isbn string) (Book, error) {
	var book Book

	row := db.handler.QueryRow("SELECT * FROM book WHERE isbn = ?", isbn)
	if err := row.Scan(&book.ISBN, &book.Title, &book.Author, &book.Summary, &book.PubYear); err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("BookById %v: no such album", isbn)
		}
		return book, fmt.Errorf("BookById %v: %v", isbn, err)
	}
	return book, nil
}

func AddBook(book Book) (int64, error) {
	result, err := db.handler.Exec("INSERT INTO book VALUES (?, ?, ?, ?, ?)",
		book.ISBN, book.Title, book.Author, book.Summary, book.PubYear)
	if err != nil {
		return 0, fmt.Errorf("addBook: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addBook: %v", err)
	}
	return id, nil
}

func DeleteBook(isbn string) (int64, error) {
	result, err := db.handler.Exec("DELETE FROM book WHERE isbn = ?", isbn)
	if err != nil {
		return 0, fmt.Errorf("addBook: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addBook: %v", err)
	}
	return id, nil
}

func GetAuthors() ([]string, error) {
	var authors []string

	rows, err := db.handler.Query("SELECT DISTINCT author FROM book")
	if err != nil {
		return nil, fmt.Errorf("GetBooks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var author string
		if err := rows.Scan(&author); err != nil {
			return nil, fmt.Errorf("GetBooks: %v", err)
		}
		authors = append(authors, author)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBooks: %v", err)
	}
	return authors, nil
}

func GetAuthorByName(name string) ([]Book, error) {
	var books []Book

	rows, err := db.handler.Query("SELECT * FROM book where author = ?", name)
	if err != nil {
		return nil, fmt.Errorf("GetBooks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ISBN, &book.Title, &book.Author, &book.Summary, &book.PubYear); err != nil {
			return nil, fmt.Errorf("GetBooks: %v", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBooks: %v", err)
	}
	return books, nil
}
