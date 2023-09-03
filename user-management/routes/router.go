package routes

import (
	"quickcooks/user-management/context"

	"github.com/gin-gonic/gin"
)

func NewRouter(context *context.UserManagementContext) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}
