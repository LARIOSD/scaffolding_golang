package boostrap

import (
	"fmt"
	"genesis/pos/presentation/middleware"

	"genesis/pos/presentation/routers"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"os"
	"strconv"
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
	port := os.Getenv("PORT")
	engine := loadEngine()

	production, _ := getBool("PRODUCTION")
	if production {
		certFile := os.Getenv("CERTFILE")
		keyFile := os.Getenv("KEYFILE")
		err := engine.RunTLS(port, certFile, keyFile)
		if err != nil {
			fmt.Println("ERROR AL INICIAR EL SERVIDOR")
			panic("ERROR AL INICIAR EL SERVIDOR:" + err.Error())
		}
	} else {
		err := engine.Run(port)
		if err != nil {
			fmt.Println("ERROR AL INICIAR EL SERVIDOR")
			panic("ERROR AL INICIAR EL SERVIDOR:" + err.Error())
		}
		fmt.Println("Servidor ejecutándose en modo desarrollo en", port)
	}
}

func getBool(key string) (bool, error) {
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return false, fmt.Errorf("error convirtiendo %s a booleano: %v", key, err)
	}
	return value, nil
}
