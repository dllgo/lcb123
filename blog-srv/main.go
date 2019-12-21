package main

import (
	"lcb123/pkg/config"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	"lcb123/blog-srv/handler"
	blog "lcb123/blog-srv/proto/blog"
	"lcb123/blog-srv/subscriber"
)

func main() {
	/************************************/
	/********** 服务发现  cousul   ********/
	/************************************/
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			config.C.Etcd,
		}
	})
	// New Service
	service := grpc.NewService(
		micro.Name(config.C.Service.Name),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*15),      //重新注册时间
		micro.RegisterInterval(time.Second*10), //注册过期时间
		micro.Version(config.C.Service.Version),
	)

	// Initialise service
	service.Init()

	// Register Handler
	blog.RegisterBlogHandler(service.Server(), new(handler.Blog))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(config.C.Service.Name, service.Server(), new(subscriber.Blog))

	// Register Function as Subscriber
	micro.RegisterSubscriber(config.C.Service.Name, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
