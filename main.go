package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})

	router.Run(":8080")
}
