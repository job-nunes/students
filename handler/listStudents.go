package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListStudentsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List students",
	})
}
