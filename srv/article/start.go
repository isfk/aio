package article

import (
	"github.com/isfk/aio/config"
	"github.com/isfk/aio/srv/article/handler"
	"github.com/isfk/aio/srv/article/model"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"github.com/isfk/aio/proto/article"
)

var configFile string

func Start(config string) {
	configFile = config
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.article"),
		micro.Version("latest"),
		micro.Action(action),
	)

	// Initialise service
	service.Init()

	// Register Handler
	article.RegisterServiceHandler(service.Server(), new(handler.Article))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	config.Init(configFile)
	model.Init(config.Conf)
	return nil
}
