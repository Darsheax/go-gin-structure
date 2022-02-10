package context

import (
	"context"
	"reflect"
	"root/Core/envRead"
	"root/Core/mailer"
	"root/context/blog"
	"root/core/auth"
	"root/core/global"
	"root/core/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Start() {

	//READ .env file
	fileEnvToRead := "context.env"
	envContext := envRead.Read(fileEnvToRead)

	env := ".env"
	envVariables := envRead.Read(env)

	//Context Initialisation
	ctx := context.TODO()

	//Context connexion: for timeout reponse
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()

	//GET connected to DB client
	DataBase := model.Connexion(ctx, envVariables["DB_NAME"], envVariables["DB_URI"])

	engine := gin.Default()

	//change f.Field() validator.ValidationErrors => return the `json tag` instead of `name` in struct
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	//Mailer
	Mailer := &mailer.Mailer{}
	Mailer.Start()

	//Authentificator
	middleware := auth.Init(DataBase, ctx, engine, Mailer)

	//DEFINE Global struct
	Global := &global.Global{
		Engine:     engine,
		DataBase:   DataBase,
		Auth:       middleware,
		AppContext: ctx,
		Mailer:     Mailer,
	}

	Global.AddContext(blog.Init())
	Global.InitContexts(envContext)

	engine.Run(":3000")
}
