package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter 
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Health check 
	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}