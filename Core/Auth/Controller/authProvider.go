package authController

import (
	"encoding/json"
	"net/http"
	authEntity "root/Core/Auth/Entity"
	authModel "root/Core/Auth/Model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func AuthProvider(middleware *jwt.GinJWTMiddleware, AuthEntity *authModel.AuthEntity, authRoute *gin.RouterGroup) {

	goth.UseProviders(
		google.New("463008552495-v343bl24nekr11utnftg9u6neuuoge45.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile"),
	)

	authRoute.GET("/:provider", func(c *gin.Context) {
		provider := c.Param("provider")
		q := c.Request.URL.Query()
		q.Add("provider", provider)
		c.Request.URL.RawQuery = q.Encode()
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})

	authRoute.GET("/:provider/callback", func(c *gin.Context) {
		provider := c.Param("provider")
		q := c.Request.URL.Query()
		q.Add("provider", provider)
		c.Request.URL.RawQuery = q.Encode()
		data, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		var user authEntity.User
		res, err := json.Marshal(data)
		err = json.Unmarshal([]byte(res), &user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		user.UseProvider = true

		/*TODO:
		If exist :
			Login
		else
			Register
		*/

		_, err = AuthEntity.Register(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"error":   err.Error(),
			})
			return
		}

		tokenString, _, err := middleware.TokenGenerator(&user)
		if err != nil {
			panic(err)
		}

		expireCookie := middleware.TimeFunc().Add(middleware.CookieMaxAge)
		maxage := int(expireCookie.Unix() - middleware.TimeFunc().Unix())

		if middleware.SendCookie {

			if middleware.CookieSameSite != 0 {
				c.SetSameSite(middleware.CookieSameSite)
			}

			c.SetCookie(
				middleware.CookieName,
				tokenString,
				maxage,
				"/",
				middleware.CookieDomain,
				middleware.SecureCookie,
				middleware.CookieHTTPOnly,
			)
		}

		c.JSON(200, gin.H{
			"message": "success",
			"user":    user,
		})

	})
}
