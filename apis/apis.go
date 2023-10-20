package server

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}


// album represents data about a record album.
type log struct {
	ID     int  `json:"id"`
	Description  string  `json:"description"`
	Commentable_type  string  `json:"commentable_type"`
	Commentable_id	int `json:"commentable_id"`
	Created_at	string `json:"created_at"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Anita", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func Get_complete_channel_list(c *gin.Context) {

	logs := GetAllChannels()
	c.IndentedJSON(http.StatusOK, logs)

}

