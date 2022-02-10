package authController

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AuthLogin(middleware *jwt.GinJWTMiddleware, authRoute *gin.RouterGroup) {
	authRoute.GET("/login", middleware.LoginHandler)
}
