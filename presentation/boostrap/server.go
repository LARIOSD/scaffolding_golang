package boostrap

import (
	"genesis/pos/presentation/middleware"
	"genesis/pos/presentation/routers"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"log"
	"os"
	"time"
)

func loadEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.RecoveryMiddleware())
	engine.Use(middleware.LoggerMiddleware())
	engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          300 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	generalGroup := engine.Group("/api")

	// customer routers
	billingGroup := generalGroup.Group("/current-time")

	routers.CurrentTime(billingGroup)

	return engine
}
func RunServer() {
	engine := loadEngine()
	engine.Use(gin.Logger())
	err := engine.Run(os.Getenv("SERVER_PORT"))
	// err := engine.Run(constants.HOST_PORT)
	if err != nil {
		log.Println("ERROR AL INICIAR EL SERVIDOR")
		panic("ERROR AL INICIAR EL SERVIDOR:" + err.Error())
	}
}
