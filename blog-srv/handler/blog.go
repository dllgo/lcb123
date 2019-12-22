package handler

import (
	"context"
	"lcb123/pkg/micros"
	pbu "lcb123/user-srv/proto/user"

	"github.com/micro/go-micro/util/log"

	blog "lcb123/blog-srv/proto/blog"
)

type Blog struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Blog) Call(ctx context.Context, req *blog.Request, rsp *blog.Response) error {
	log.Log("Received Blog.Call request")
	userClient := pbu.NewUserService("com.lcb123.srv.user", micros.GetService().Client())
	rspu, _ := userClient.Call(ctx, &pbu.Request{
		Name: "67890",
	})
	rsp.Msg = rspu.Msg
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Blog) Stream(ctx context.Context, req *blog.StreamingRequest, stream blog.Blog_StreamStream) error {
	log.Logf("Received Blog.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&blog.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Blog) PingPong(ctx context.Context, stream blog.Blog_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&blog.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
