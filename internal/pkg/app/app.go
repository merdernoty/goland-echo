package app

import (
	"example.com/goland-echo/internal/app/controller"
	"example.com/goland-echo/internal/app/mw"
	"example.com/goland-echo/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type App struct {
	c    *controller.Controller
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {

	a := &App{}

	a.s = service.New()

	a.c = controller.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.UserCheck)

	a.echo.GET("/status", a.c.Status)

	return a, nil
}

func (a *App) Run() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("starting server")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("could not start server")
	}
	return nil
}
