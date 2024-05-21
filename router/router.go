package router

import (
	"garrison/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", api.GetUsers)
	router.GET("/users/:id", api.GetUserById)
	router.POST("/user", api.CreateUser)

	return router
}
