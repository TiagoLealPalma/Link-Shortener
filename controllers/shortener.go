package controllers

import (
	"Link-Shortener/data"
	"Link-Shortener/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShortenURL(c *gin.Context) {
	var json struct {
		OriginalLink string `json:"url"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Printf("json binding failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ogLink := "original Link: " + json.OriginalLink
	print(ogLink)

	// Generate Token
	token := utils.GenerateShortCode(6)
	shortUrl := "http://localhost:8080/" + token

	// Save to DataBase
	err := data.SaveUrl(token, json.OriginalLink)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortUrl})
	print(shortUrl)

}

func print(str string) {
	fmt.Printf("===========DEBUG===========\n")
	fmt.Printf(str + "\n")
	fmt.Printf("===========DEBUG===========\n")
}
