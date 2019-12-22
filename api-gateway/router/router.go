package router

import (
	"lcb123/pkg/log"
	"lcb123/pkg/trace"

	"github.com/gin-gonic/gin"

	"lcb123/api-gateway/handler"
	"lcb123/api-gateway/middleware"
)

func Init() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) //是否生产模式启动
	// validator.Init()             //自定义验证器 替换 gin 里面的验证器
	router := gin.Default()

	router.NoRoute(middleware.NoRouteHandler())
	// 崩溃恢复
	router.Use(middleware.RecoveryMiddleware())
	// gin日志
	router.Use(log.GinLogger())

	// jaeger trace 追踪
	router.Use(trace.TracerWrapper)
	// 跨域
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")                   //跨域
		ctx.Header("Access-Control-Allow-Headers", "token,Content-Type") //必须的请求头
		ctx.Header("Access-Control-Allow-Methods", "OPTIONS,POST,GET")   //接收的请求方法
	})
	r := router.Group("/user")
	// use the generated client stub
	r.GET("/test", handler.Anything)
	r.GET("/test1", handler.Anything1)
	r.GET("/test2", handler.Anything2)
	return router
}
