package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/trananhtung/url-shortener/config"
	"github.com/trananhtung/url-shortener/models"
	"github.com/trananhtung/url-shortener/orm"
)

var link = config.SITE{}
var DB = orm.DB{}

func init() {
	fmt.Println("Initializing config...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	link.Init()
	DB.Init()

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

			err = DB.Create(json.HashURL(), json.URL)

			if err != nil {
				fmt.Println(err.Error())
			}

			shortURL := url + "/" + json.HashURL()
			c.JSON(http.StatusOK, gin.H{"url": shortURL})
		}
	})

	r.POST("/shorten/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		var json models.CreateShortenURL

		if c.Bind(&json) == nil {
			err := json.Validate()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			err = DB.Create(hash, json.URL)

			if err != nil {
				fmt.Println(err.Error())
			}

			shortURL := url + "/" + hash

			c.JSON(http.StatusOK, gin.H{"url": shortURL})
		}
	})

	r.GET("/all", func(c *gin.Context) {
		urls, err := DB.FindAll(url)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"urls": urls})
	})

	r.GET("/record", func(c *gin.Context) {
		urls, err := DB.FindAllRecordIP()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"urls": urls})
	})

	r.GET("/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		url, err := DB.Find(hash)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		go DB.CreateRecordIP(hash, c.ClientIP())

		c.Redirect(http.StatusMovedPermanently, url)
	})

	r.Run(":8080")
}
