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
		c.Redirect(http.StatusMovedPermanently, "http://www.magdaleal.pt")
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
