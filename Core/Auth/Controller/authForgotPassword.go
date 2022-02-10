package authController

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	authEntity "root/Core/Auth/Entity"
	"root/Core/mailer"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AuthForgotPassword(middleware *jwt.GinJWTMiddleware, mailer *mailer.Mailer, authRoute *gin.RouterGroup) {
	authRoute.GET("/forgot/:email", func(c *gin.Context) {

		userEmail := c.Param("email")
		token, _, _ := middleware.TokenGenerator(&authEntity.User{
			Email: userEmail,
		})

		wd, _ := os.Getwd()

		mailTemplate, err := template.ParseFiles(wd + "/Core/mailer/Template/forgotPassword.html")
		if err != nil {
			panic(err)
		}

		var body bytes.Buffer

		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

		to := []string{
			userEmail,
		}

		mailTemplate.Execute(&body, struct {
			Link string
		}{
			Link: c.Request.Host + "/reset/" + token, // HTTPS !
		})

		// Sending email.
		mailer.Send(to, body.Bytes())

		c.JSON(200, gin.H{
			"token": token,
		})

	})
}
