package main

import "github.com/gin-gonic/gin"
import "github.com/Junbong/coinsightly-api/services"

func main() {
	// Run background fetcher
	go services.RunFetcher()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
