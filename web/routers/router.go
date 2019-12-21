package routers

import (
	"lcb123/pkg/log"
	"lcb123/pkg/trace"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //是否生产模式启动
	router := gin.Default()

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

	return router
}
