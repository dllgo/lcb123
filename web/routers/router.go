package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 用户登陆接口
// @Summary 用户登陆接口
// @Tags UserControl
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param code     query string false "验证码"
// @Success 200 {object} response.JsonObject
// @Router /user/api/login [post]

// 路由总入口，注册所有微服务的 路由
func Register(router *gin.Engine) {
	//配置跨域
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "ACCESS_TOKEN"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	router.HandleMethodNotAllowed = true
	// 使用gin-swagger 中间件
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
