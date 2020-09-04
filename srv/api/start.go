package api

import (
	"github.com/isfk/aio/config"
	"github.com/isfk/aio/srv/api/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"net/http"
)

var configFile string

func Start(config string) {
	configFile = config
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.api"),
		web.Version("latest"),
		web.Action(action),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	//register handler
	service.Handle("/", echoInit())

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func echoInit() *echo.Echo {
	// Echo instance
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	return router.Router(e)
}

func action(c *cli.Context) {
	config.Init(configFile)
}
