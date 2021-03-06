package http

import (
	"github.com/go-colon/colon/app/http/module/demo"
	"github.com/go-colon/colon/core/contract"
	"github.com/go-colon/colon/core/gin"
	ginSwagger "github.com/go-colon/colon/core/middleware/gin-swagger"
	"github.com/go-colon/colon/core/middleware/gin-swagger/swaggerFiles"
	"github.com/go-colon/colon/core/middleware/static"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {
	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//外层定义路由
	r.GET("/hello", demo.NewDemoApi().DemoHello)

	// 动态路由定义
	demo.Register(r)
}
