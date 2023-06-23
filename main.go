package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/trananhtung/url-shortener/config"
	"github.com/trananhtung/url-shortener/models"
)

var link = config.SITE{}

func init() {
	fmt.Println("Initializing config...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	link.Init()
}

func main() {
	r := gin.Default()

	url := link.GetSite()

	r.POST("/shorten", func(c *gin.Context) {
		var json models.CreateShortenURL

		if c.Bind(&json) == nil {
			err := json.Validate()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			fmt.Println(json)

			shortURL := url + "/" + json.HashURL()
			c.JSON(http.StatusOK, gin.H{"url": shortURL})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
