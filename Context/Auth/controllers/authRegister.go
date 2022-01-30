package authController

import (
	authModel "root/Context/Auth/Models"
	global "root/Core/Global"
	user "root/Core/User"
	utility "root/Core/Utility"

	"github.com/gin-gonic/gin"
)

func AuthRegister(global *global.Global, auth *gin.RouterGroup) {
	auth.GET("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"route": "/blog/comment",
		})
	})

	name := "Bastien"
	pswd := "bb"

	password, _ := utility.HashPassword(pswd)

	user := user.User{
		Name:     name,
		Password: password,
	}

	registerEntity := authModel.AuthRegisterEntity{
		DataBase:   global.DataBase,
		AppContext: global.AppContext,
	}

	registerEntity.Register(user)
}