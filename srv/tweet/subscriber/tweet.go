package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	tweet "github.com/isfk/aio/srv/tweet/proto/tweet"
)

type Tweet struct{}

func (e *Tweet) Handle(ctx context.Context, msg *tweet.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *tweet.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
