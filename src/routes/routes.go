package routes

import (
	"github.com/JamestarKurbah/go-auth-app/src/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)
	router.GET("/token", handlers.TokenHandler)

	return router
}
