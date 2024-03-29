package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello-world",
		})
	})
	fmt.Println("Server has been started")
	r.Run() // listen and serve on 0.0.0.0:8080
}
