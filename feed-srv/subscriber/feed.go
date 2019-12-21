package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	feed "lcb123/feed-srv/proto/feed"
)

type Feed struct{}

func (e *Feed) Handle(ctx context.Context, msg *feed.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *feed.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
