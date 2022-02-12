package blogController

import (
	"root/Core/global"
	"root/core/auth/authEntity"
	"root/core/auth/authMiddleware"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func BlogComment(global *global.Global) {
	r := global.Engine.Group("/blog/comment")

	r.Use(global.Auth.MiddlewareFunc())

	r.GET("/", func(c *gin.Context) {
		user, _ := c.Get(authMiddleware.IdentityKey)

		localizer := global.Translator.Localizer(c)

		message := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "BlogComment",
				Other: "Bonjour {{.Name}}, bienvenue sur le blog",
			},
			TemplateData: map[string]string{
				"Name": user.(*authEntity.User).Name,
			},
		})

		c.JSON(200, gin.H{
			"userID":  user.(*authEntity.User).Email,
			"message": message,
		})
	})
}
