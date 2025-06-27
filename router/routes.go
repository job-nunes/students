package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Define your routes here
	v1 := router.Group("/api/v1")
	{
		v1.GET("/students", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET student",
			})
		})
		v1.GET("/students", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET students",
			})
		})
		v1.POST("/students", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "POST student",
			})
		})
		v1.PUT("/students", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "PUT student",
			})
		})
		v1.DELETE("/students", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE student",
			})
		})
	}
}
