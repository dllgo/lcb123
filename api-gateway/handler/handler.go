package handler

import (
	pbb "lcb123/blog-srv/proto/blog"
	pbf "lcb123/feed-srv/proto/feed"
	"lcb123/pkg/log"
	"lcb123/pkg/micros"
	"lcb123/pkg/trace"
	pb "lcb123/user-srv/proto/user"

	"github.com/gin-gonic/gin"
)

var userClient pb.UserService
var feedClient pbf.FeedService
var blogClient pbb.BlogService

func init() {
	userClient = pb.NewUserService("com.lcb123.srv.user", micros.GetService().Client())
	feedClient = pbf.NewFeedService("com.lcb123.srv.feed", micros.GetService().Client())
	blogClient = pbb.NewBlogService("com.lcb123.srv.blog", micros.GetService().Client())
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

//测试
func Anything1(c *gin.Context) {
	ctx, ok := trace.ContextWithSpan(c)
	if !ok {
		log.Warn("不存在context")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "不存在context",
		})
		return
	}
	log.Error(feedClient)
	rsp, _ := feedClient.Call(ctx, &pbf.Request{
		Name: "45678",
	})
	log.Error(rsp)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  rsp.Msg,
	})
}

//测试
func Anything2(c *gin.Context) {
	ctx, ok := trace.ContextWithSpan(c)
	if !ok {
		log.Warn("不存在context")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "不存在context",
		})
		return
	}
	log.Error(blogClient)
	rsp, _ := blogClient.Call(ctx, &pbb.Request{
		Name: "67890",
	})
	log.Error(rsp)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  rsp.Msg,
	})
}
