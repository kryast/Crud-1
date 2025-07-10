package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-1.git/handlers"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)

	return r
}
