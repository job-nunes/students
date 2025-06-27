package router

import (
	"github.com/gin-gonic/gin"
	"github.com/job-nunes/students/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Define your routes here
	v1 := router.Group("/api/v1")
	{
		v1.GET("/students/:id", handler.ShowStudentHandler)
		v1.GET("/students", handler.ListStudentsHandler)
		v1.POST("/students", handler.CreateStudentHandler)
		v1.PUT("/students", handler.UpdateStudentHandler)
		v1.DELETE("/students", handler.DeleteStudentHandler)
	}
}
