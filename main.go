package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/echo", func(c *gin.Context) {
		var message map[string]interface{}

		if err := c.BindJSON(&message); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": message})
	})

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run(":8889")
}
