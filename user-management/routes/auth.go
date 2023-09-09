package routes

import (
	"net/http"
	"quickcooks/user-management/context"

	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(rg *gin.RouterGroup, context *context.UserManagementContext) {
	authRoutes := rg.Group("/auth")

	authRoutes.POST("/login", func(c *gin.Context) {
		type LoginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var loginRequest LoginRequest

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		token, user, authed, err := context.AuthenticationService.Login(loginRequest.Email, loginRequest.Password)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if authed {
			c.SetCookie("token", token, 3600, "/", "", true, true)
			c.JSON(http.StatusOK, user)
		}
		if !authed {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	})
}
