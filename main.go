package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: "1", Title: "1984", Author: "George Orwell"},
    {ID: "2", Title: "Brave New World", Author: "Aldous Huxley"},
    {ID: "3", Title: "Fahrenheit 451", Author: "Ray Bradbury"},
}

func main() {
    router := gin.Default()

    // Получение всех книг
    router.GET("/books", getBooks)

    // Получение книги по ID
    router.GET("/books/:id", getBookByID)

    // Создание новой книги
    router.POST("/books", createBook)

    // Обновление существующей книги
    router.PUT("/books/:id", updateBook)

    // Удаление книги
    router.DELETE("/books/:id", deleteBook)

    router.Run(":8080")
}

func getBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")

    for _, book := range books {
        if book.ID == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func createBook(c *gin.Context) {
    var newBook Book

    if err := c.BindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
        return
    }

    books = append(books, newBook)
    c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook Book

    if err := c.BindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
        return
    }

    for i, book := range books {
        if book.ID == id {
            books[i] = updatedBook
            c.JSON(http.StatusOK, updatedBook)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBook(c *gin.Context) {
    id := c.Param("id")

    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}