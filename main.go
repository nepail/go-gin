package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// Hash
		c.JSON(200, gin.H{
			"message": "pongdddd",
		})
	})
	router.Run() // listen and serve on 0.0.0.0: 8080
}
