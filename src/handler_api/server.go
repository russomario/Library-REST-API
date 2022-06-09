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

package handler_api

import (
	"fmt"
	"net/http"

	"example.com/database_manager"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	database_manager.SetUpDBConnection()
	router := gin.Default()

	// Routes for Books
	bookEndpoint := router.Group("/books")
	{
		bookEndpoint.GET("", getBooks)
		bookEndpoint.GET(":isbn", getBookByID)
		bookEndpoint.POST("", postBook)
		bookEndpoint.DELETE(":isbn", deleteBookByID)
	}

	// Routes for authors
	authorEndpoint := router.Group("/authors")
	{
		authorEndpoint.GET("", getAuthors)
		authorEndpoint.GET(":name", getAuthorBookByName)
	}

	router.Run("localhost:8080")
}

func postBook(c *gin.Context) {
	var newBook database_manager.Book

	// Call BindJSON to bind the received JSON to newBook.
	if err := c.BindJSON(&newBook); err != nil {
		fmt.Printf("Type of error: %v\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "make sure to fill of required field as well as to use the correct type"})
		return
	}

	if id, err := database_manager.AddBook(newBook); err == nil {
		fmt.Println(id)
		c.IndentedJSON(http.StatusCreated, newBook)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Duplicate entry"})
	}
}

func getBooks(c *gin.Context) {
	books, err := database_manager.GetBooks()
	if err == nil {
		if len(books) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no books are stored"})
		} else {
			c.IndentedJSON(http.StatusOK, books)
		}
	} else {
		fmt.Printf("error no isbn: %v", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving data"})
	}
}

func getBookByID(c *gin.Context) {
	book, err := database_manager.GetBookByISBN(c.Param("isbn"))
	if err == nil {
		c.IndentedJSON(http.StatusOK, book)
	} else {
		fmt.Printf("error no isbn: %v", err.Error())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "There is no book with this ISBN"})
	}
}

func deleteBookByID(c *gin.Context) {
	isbn := c.Param("isbn")
	if _, err := database_manager.DeleteBook(isbn); err == nil {
		c.IndentedJSON(http.StatusNoContent, nil)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No entity"})
	}
}

func getAuthors(c *gin.Context) {
	books, err := database_manager.GetAuthors()
	if err == nil {
		if len(books) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no books are stored"})
		} else {
			c.IndentedJSON(http.StatusOK, books)
		}
	} else {
		fmt.Printf("error no isbn: %v", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving data"})
	}
}

func getAuthorBookByName(c *gin.Context) {
	author, err := database_manager.GetAuthorByName(c.Param("name"))
	if err == nil {
		c.IndentedJSON(http.StatusOK, author)
	} else {
		fmt.Printf("error no isbn: %v", err.Error())
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "There is no book with this ISBN"})
	}
}
