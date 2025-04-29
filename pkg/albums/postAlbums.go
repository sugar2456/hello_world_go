package albums

import (
	"hello_world_go/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAlbums(c *gin.Context) {
	var newAlbum types.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
