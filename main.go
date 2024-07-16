package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// book represents data about a book.
type book struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	Author            string  `json:"author"`
	YearOfPublication int     `json:"publication_year"`
	Price             float64 `json:"price"`
}

// books slice to seed record book data.
var books = []book{
	{ID: "1", Title: "Life Stories", Author: "Heather Newbold", YearOfPublication: 1998, Price: 4.99},
	{ID: "2", Title: "All Around The Town", Author: "Mary Higgins Clark", YearOfPublication: 2020, Price: 5.29},
	{ID: "3", Title: "The 48 Laws of Power", Author: "Robert Greene", YearOfPublication: 2023, Price: 10.99},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBook)

	router.Run("localhost:8080")
}

// getBooks responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBook adds a book from JSON received in the request body.
func postBook(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of books, looking for
	// an book whose ID value matches the parameter.
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
