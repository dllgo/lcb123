package main

import (
	"lcb123/pkg/config"
	"lcb123/pkg/log"
	"lcb123/pkg/micros"

	"github.com/micro/go-micro"

	"lcb123/user-srv/handler"
	user "lcb123/user-srv/proto/user"
	"lcb123/user-srv/subscriber"
)

func main() {
	service := micros.GetService()
	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))
	// Register Struct as Subscriber
	micro.RegisterSubscriber(config.C.Service.Name, service.Server(), new(subscriber.User))
	// Register Function as Subscriber
	micro.RegisterSubscriber(config.C.Service.Name, service.Server(), subscriber.Handler)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
