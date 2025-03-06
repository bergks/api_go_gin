package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Film struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Director string `json:"director"`
	Year string `json:"year"`
}

var films = []Film{
	{ID: "1", Title: "Titanic", Director: "James Cameron", Year: "1997"},
	{ID: "2", Title: "Avatar", Director: "James Cameron", Year: "2009"},
	{ID: "3", Title: "The Great Gatsby", Director: "Baz Luhrmann", Year: "2013"},
	{ID: "4", Title: "Spider-Man", Director: "Sam Raimi", Year: "2002"},
	{ID: "5", Title: "The Matrix", Director: "Lana and Lilly Wachowski", Year: "1999"},
}

func main() {
	router := gin.Default()

	// Получение всех фильмов
	router.GET("/films", getFilms)

	// Получение книги по ID
	router.GET("/films/:id", getFilmByID)

	// Создание новой книги
	router.POST("/films", createFilm)

	// Обновление существующей книги
	router.PUT("/filmss/:id", updateFilm)

	// Удаление книги
	router.DELETE("/films/:id", deleteFilm)

	router.Run(":8080")
}

func getFilms(c *gin.Context) {
	c.JSON(http.StatusOK, films)
}

func getFilmByID(c *gin.Context) {
	id := c.Param("id")

	for _, film := range films {
		if film.ID == id {
			c.JSON(http.StatusOK, film)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "film not found"})
}

func createFilm(c *gin.Context) {
	var newFilm Film

	if err := c.BindJSON(&newFilm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	films = append(films, newFilm)
	c.JSON(http.StatusCreated, newFilm)
}

func updateFilm(c *gin.Context) {
	id := c.Param("id")
	var updatedFilm Film

	if err := c.BindJSON(&updatedFilm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, book := range films {
		if book.ID == id {
			films[i] = updatedFilm
			c.JSON(http.StatusOK, updatedFilm)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteFilm(c *gin.Context) {
	id := c.Param("id")

	for i, film := range films {
		if film.ID == id {
			films = append(films[:i], films[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "film deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "film not found"})
}
