package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// estructura de datos REST API de Album
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// creamos la data
var albums = []album{
	{ID: "1", Title: "ambulancia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "camisa negra", Artist: "Juanes", Year: 2023},
	{ID: "3", Title: "maquina", Artist: "Anuel", Year: 2020},
}

// obtener lista de albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// utilizando post
func postAlbums(c *gin.Context) {
	var newAlbum album
	/*c.BindJSON(&newAlbum)
	--- manejo de ERROR*/
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

// controlador DV ART
func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

// funcion principal
func main() {
	//fmt.Println("Hola mundo")
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.POST("/albums", postAlbums)

	router.GET("/albums/:id", getAlbumsByID)

	router.Run("localhost:8080")
}
