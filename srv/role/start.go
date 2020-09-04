package role

import (
	"github.com/isfk/aio/config"
	role "github.com/isfk/aio/proto/role"
	"github.com/isfk/aio/srv/role/handler"
	"github.com/isfk/aio/srv/role/model"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

var configFile string

func Start(config string) {
	configFile = config
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.role"),
		micro.Version("latest"),
		micro.Action(action),
	)

	// Initialise service
	service.Init()

	// Register Handler
	role.RegisterServiceHandler(service.Server(), new(handler.Role))

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
