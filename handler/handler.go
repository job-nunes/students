package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowStudentHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET student",
	})
}
func ListStudentsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List students",
	})
}
func CreateStudentHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST student",
	})
}
func UpdateStudentHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT student",
	})
}
func DeleteStudentHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE student",
	})
}
