package controllers

import (
	"Link-Shortener/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect(c *gin.Context) {
	token := c.Param("token")

	// Access the database and get the link mapped to the token
	url, err := data.GetOriginalLink(token)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
