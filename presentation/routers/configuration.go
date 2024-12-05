package routers

import (
	"genesis/pos/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func CurrentTime(router *gin.RouterGroup) {
	router.GET("/", handlers.GetCurrentTime)
}
