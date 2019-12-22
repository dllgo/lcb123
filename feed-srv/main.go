package main

import (
	"lcb123/pkg/config"
	"lcb123/pkg/log"
	"lcb123/pkg/micros"

	"github.com/micro/go-micro"

	"lcb123/feed-srv/handler"
	feed "lcb123/feed-srv/proto/feed"
	"lcb123/feed-srv/subscriber"
)

func main() {

	// Initialise service
	service := micros.GetService()

	// Register Handler
	feed.RegisterFeedHandler(service.Server(), new(handler.Feed))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(config.C.Service.Name, service.Server(), new(subscriber.Feed))

	// Register Function as Subscriber
	micro.RegisterSubscriber(config.C.Service.Name, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
