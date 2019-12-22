package handler

import (
	"lcb123/pkg/log"
	"lcb123/pkg/micros"
	"lcb123/pkg/trace"
	pb "lcb123/user-srv/proto/user"

	"github.com/gin-gonic/gin"
)

var userClient pb.UserService

func init() {
	userClient = pb.NewUserService("com.lcb123.srv.user", micros.GetService().Client())
}

//测试
func Anything(c *gin.Context) {
	ctx, ok := trace.ContextWithSpan(c)
	if !ok {
		log.Warn("不存在context")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "不存在context",
		})
		return
	}
	log.Error(userClient)
	rsp, _ := userClient.UserDetail(ctx, &pb.UserRequest{
		Uid: "1",
	})
	log.Error(rsp)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  rsp,
	})
}
