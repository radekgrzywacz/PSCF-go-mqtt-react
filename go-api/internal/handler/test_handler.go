package handler

import "github.com/gin-gonic/gin"

// Test handles GET /api/test and returns "Hello World"
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
