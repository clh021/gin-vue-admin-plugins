package study_plugin

import (
	"github.com/gin-gonic/gin"
)

var StudyPlugin = new(pluginStudy)

type pluginStudy struct{}

func (*pluginStudy) RouterPath() string {
	return "study"
}

func (*pluginStudy) Register(group *gin.RouterGroup) {
	//如需细分权限 可以在此处use中间件 gva项目包名已改为github模式
	//所以整个plugin可以直接独立到外层开启为新的项目 然后用包的形式导入也是可以完整运行的
	// 例:
	/*
		group.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).GET("hello", func(context *gin.Context) {
			context.JSON(200, "hello world")
		})
	*/
	group.GET("hello", func(context *gin.Context) {
		context.JSON(200, "hello world")
	})
	group.GET("1", s1)
}
