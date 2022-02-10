package authController

import (
	"fmt"
	"net/http"
	"net/url"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AuthRefreshToken(middleware *jwt.GinJWTMiddleware, authRoute *gin.RouterGroup) {

	authRoute.GET("/refresh_token/*referer", func(c *gin.Context) {
		if _, _, err := middleware.RefreshToken(c); err != nil {
			fmt.Println(err)
			login := url.URL{Path: "/login"}
			c.Redirect(http.StatusFound, login.RequestURI())
		} else {
			referer := url.URL{Path: c.Param("referer")}
			c.Redirect(http.StatusFound, referer.RequestURI())
		}
	})

}
