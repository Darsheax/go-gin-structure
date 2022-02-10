package authMiddleware

import (
	"root/core/auth/authEntity"
	"root/core/auth/authModel"
	"root/core/utility"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

var (
	key         = []byte("secret key")
	IdentityKey = "mail"

	timeout        = time.Minute * 20
	cookieMaxAge   = time.Hour * 24 * 30
	sendCookie     = true
	secureCookie   = false
	cookieHTTPOnly = true
)

func Middleware(AuthEntity *authModel.AuthEntity) *jwt.GinJWTMiddleware {

	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:        "Authentificator",
		Key:          key,
		Timeout:      timeout,
		MaxRefresh:   cookieMaxAge,
		CookieMaxAge: cookieMaxAge,
		IdentityKey:  IdentityKey,

		SendCookie:     sendCookie,
		SecureCookie:   secureCookie, // true when in prod with https
		CookieHTTPOnly: cookieHTTPOnly,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*authEntity.User); ok {
				return jwt.MapClaims{
					IdentityKey: user.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &authEntity.User{
				Email: claims[IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals authEntity.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user := AuthEntity.AuthLogin(loginVals.Email)
			if (authEntity.User{} != user && utility.CheckPasswordHash(loginVals.Password, user.Password)) {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*authEntity.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(200, gin.H{
				"code":    code,
				"message": message,
			})
			//refresh := url.URL{Path: "/auth/refresh_token/" + c.Request.URL.Path}
			//c.Redirect(http.StatusFound, refresh.RequestURI())
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	return authMiddleware

}
