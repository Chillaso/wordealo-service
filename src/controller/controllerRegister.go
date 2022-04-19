package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterController(router *gin.Engine) {
	router.POST("/wordealo", GetWordle)
}
