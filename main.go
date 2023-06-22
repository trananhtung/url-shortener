package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/trananhtung/url-shortener/models"
	"github.com/trananhtung/url-shortener/utils"
)

func init() {
	fmt.Println("Initializing config...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func main() {
	r := gin.Default()
	link := models.URL{}

	link.Init()
	fmt.Println(link.GetURL())
	url := link.GetURL()

	r.POST("/shorten", func(c *gin.Context) {
		var json struct {
			URL string `json:"url" binding:"required"`
		}

		if c.Bind(&json) == nil {

			// Validate the URL
			err := utils.ValidateURL(json.URL)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			shortURL := url + "/" + utils.HashURL(json.URL)
			c.JSON(http.StatusOK, gin.H{"shortened": shortURL})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
