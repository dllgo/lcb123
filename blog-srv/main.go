package main

import (
	"lcb123/pkg/config"
	"lcb123/pkg/log"
	"lcb123/pkg/micros"

	"github.com/micro/go-micro"

	"lcb123/blog-srv/handler"
	blog "lcb123/blog-srv/proto/blog"
	"lcb123/blog-srv/subscriber"
)

func main() {

	// Initialise service
	service := micros.GetService()
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
