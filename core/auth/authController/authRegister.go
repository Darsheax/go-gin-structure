package authController

import (
	"net/http"
	"root/core/auth/authEntity"
	"root/core/auth/authModel"
	"root/core/utility"

	"github.com/gin-gonic/gin"
)

func AuthRegister(AuthEntity *authModel.AuthEntity, authRoute *gin.RouterGroup) {
	authRoute.GET("/register", func(c *gin.Context) {

		var register authEntity.Register
		if err := c.ShouldBind(&register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utility.BeautifulError(err)})
			return
		}

		password, _ := utility.HashPassword(register.Password)

		user := authEntity.User{
			Name:     register.UserName,
			Email:    register.Email,
			Password: password,
		}

		id, err := AuthEntity.Register(user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message":  "success",
			"username": user.Name,
			"email":    user.Email,
			"password": user.Password,
			"ID":       id,
		})

	})
}
