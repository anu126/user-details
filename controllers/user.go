package controllers

import (
	// Import the Gin library
	"github.com/gin-gonic/gin"
)


// HelloWorldController handles returning the hello world JSON response
func HelloWorldController(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello world, climate change is real"})
}
