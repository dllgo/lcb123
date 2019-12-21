package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	feed "lcb123/feed-srv/proto/feed"
)

type Feed struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Feed) Call(ctx context.Context, req *feed.Request, rsp *feed.Response) error {
	log.Log("Received Feed.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Feed) Stream(ctx context.Context, req *feed.StreamingRequest, stream feed.Feed_StreamStream) error {
	log.Logf("Received Feed.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&feed.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Feed) PingPong(ctx context.Context, stream feed.Feed_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&feed.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
