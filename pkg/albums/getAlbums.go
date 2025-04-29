package albums

import (
	"hello_world_go/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// パッケージレベルの定義なので、同パッケージ内の関数からアクセスすることができる
var albums = []types.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
