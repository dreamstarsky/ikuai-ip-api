package main

import (
	"ikuai-ip-api/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	r := gin.Default()

	url := os.Getenv("URL")
	username := os.Getenv("User")
	password := os.Getenv("PASSWORD")

	if url == "" || username == "" || password == "" {
		log.Println("$URL, $User or $PASSWORD not found!")
		os.Exit(1)
	}

	ikuai := api.NewIkuai(url, username, password)

	err := ikuai.Login()
	if err != nil {
		log.Println("Initial login failed:", err)
	} else {
		log.Println("Initial login successful")
	}

	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		
		c.JSON(200, gin.H{
			"dawd": 114,
			"id":   id,
		})
	})

	r.Run(":8080")
}
