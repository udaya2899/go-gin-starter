package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udaya2899/go-gin-starter/storage"
)

// New returns a new instance of a gin server
func New(repository *storage.Repository) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", getStatus)

	return r
}

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
