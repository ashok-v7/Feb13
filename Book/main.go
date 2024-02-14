package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)

// GetBooks responds with the list of all books as JSON.
func GetBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

// CreateBook adds a new book from JSON received in the request body.
func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusCreated, book)
}

// UpdateBook updates the details of a book identified by its ID.
func UpdateBook(c *gin.Context) {
	// Get ID from the request URL
	var book Book
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	// Bind the request body to the book instance
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&book)
	c.JSON(http.StatusOK, book)
}

// DeleteBook removes a book identified by its ID.
func DeleteBook(c *gin.Context) {
	// Get ID from the request URL
	var book Book
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book has been deleted"})
}

func main() {
	db = ConnectDatabase()
	db.AutoMigrate(&Book{})

	r := gin.Default()

	r.GET("/books", GetBooks)
	r.POST("/books", CreateBook)
	r.PUT("/books/:id", UpdateBook)    // Route for updating a book
	r.DELETE("/books/:id", DeleteBook) // Route for deleting a book

	r.Run() // By default, it listens on :8080
}
