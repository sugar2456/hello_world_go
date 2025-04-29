package main

import (
	"hello_world_go/pkg/albums"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", albums.GetAlbums)

	router.Run("localhost:8080")
}
