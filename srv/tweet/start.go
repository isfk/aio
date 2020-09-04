package tweet

import (
	"github.com/isfk/aio/srv/tweet/handler"
	"github.com/isfk/aio/srv/tweet/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	tweet "github.com/isfk/aio/proto/tweet"
)

func Start() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.tweet"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	tweet.RegisterTweetHandler(service.Server(), new(handler.Tweet))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.tweet", service.Server(), new(subscriber.Tweet))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
