package router

import (
	"hello/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/hello", controllers.Hello)

	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Server Failed To Connect")
	}
}
