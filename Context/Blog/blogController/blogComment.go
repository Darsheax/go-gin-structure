package blogController

import (
	"root/Core/global"
	"root/core/auth/authEntity"
	"root/core/auth/authMiddleware"

	"github.com/gin-gonic/gin"
)

func BlogComment(global *global.Global) {

	r := global.Engine.Group("/blog/comment")

	r.Use(global.Auth.MiddlewareFunc())

	r.GET("/", func(c *gin.Context) {
		user, _ := c.Get(authMiddleware.IdentityKey)

		c.JSON(200, gin.H{
			"userID": user.(*authEntity.User).Email,
		})
	})
}
