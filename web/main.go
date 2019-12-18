package main

import (
	"lcb123/web/routers"

	_ "lcb123/web/docs"
	"github.com/gin-gonic/gin"
)

// @title Swagger Example API12222
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1

func main() {
	router := gin.Default()
	routers.Register(router)
	// 指定地址和端口号
	router.Run("localhost:9090")

}
