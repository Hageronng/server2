package controllers

import (
	"github.com/gin-gonic/gin"
)

func PostUserData(ctx *gin.Context, email string, password string) {
	ctx.JSON(200, gin.H{
		"Email":    email,
		"Password": password,
	})
}
