package blog

import (
	"root/Core/global"
	"root/context/blog/blogController"
)

func Init() *global.ContextController {

	context := &global.ContextController{
		Name: "blog",
		Start: func(global *global.Global) {
			blogController.BlogUser(global)
			blogController.BlogComment(global)
		},
	}

	return context
}
