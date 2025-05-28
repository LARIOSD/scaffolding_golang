package routers

import (
	"genesis/pos/presentation/handlers"
	"genesis/pos/presentation/middleware"
	"genesis/pos/presentation/schema"
	"github.com/gin-gonic/gin"
)

func CurrentTime(router *gin.RouterGroup) {
	// example
	router.POST("/:id",
		middleware.ValidateParams(new(schema.CurrentTimeParams)),
		middleware.ValidateQuery(new(schema.CurrentTimeQuery)),
		middleware.ValidateBody(new(schema.CurrentTimeBody)),
		handlers.GetCurrentTime)
}
