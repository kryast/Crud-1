package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-1.git/handlers"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUser)

	return r
}
