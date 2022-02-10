package authController

import (
	"net/http"
	"root/core/auth/authEntity"
	"root/core/auth/authModel"
	"root/core/utility"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AuthResetPassword(middleware *jwt.GinJWTMiddleware, AuthEntity *authModel.AuthEntity, authRoute *gin.RouterGroup) {

	authRoute.POST("/reset", func(c *gin.Context) {

		params := c.Request.URL.Query()

		if !params.Has("token") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token not found",
			})
			return
		}

		token := params.Get("token")

		tokenParsed, err := middleware.ParseTokenString(token)
		if !tokenParsed.Valid || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token invalid",
			})
			return
		}

		//TODO: create a binding <password>
		var resetPassword authEntity.ResetPassword
		if err := c.ShouldBind(&resetPassword); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utility.BeautifulError(err)})
			return
		}

		if resetPassword.Password != resetPassword.PasswordVerify {
			c.JSON(http.StatusUnauthorized, gin.H{
				"errors": "Les mots de passe doivent etre identiques",
			})
			return
		}

		tokenExtract := jwt.ExtractClaimsFromToken(tokenParsed)
		userMail := tokenExtract["mail"].(string)

		password, _ := utility.HashPassword(resetPassword.Password)

		if err := AuthEntity.ResetPassword(userMail, password); err != nil {
			c.JSON(200, gin.H{
				"errors": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})
}
