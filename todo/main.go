package main

import "github.com/gin-gonic/gin"

func main() {
	const PORT = "8009"
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":" + PORT)
}
