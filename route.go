package main

import (
	"Hade/framework"
	"Hade/framework/middleware"
)

func registerRouter(core *framework.Core) {
	// 需求1+2:HTTP方法+静态路由匹配
	core.Get("/user/login", middleware.Test2(), UserLoginController)
	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test2())
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
