package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	tweet "github.com/isfk/aio/srv/tweet/proto/tweet"
)

type Tweet struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Tweet) Call(ctx context.Context, req *tweet.Request, rsp *tweet.Response) error {
	log.Info("Received Tweet.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Tweet) Stream(ctx context.Context, req *tweet.StreamingRequest, stream tweet.Tweet_StreamStream) error {
	log.Infof("Received Tweet.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&tweet.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Tweet) PingPong(ctx context.Context, stream tweet.Tweet_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&tweet.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
