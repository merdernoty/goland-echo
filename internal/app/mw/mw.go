package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strings"
)

const userAdmin = "igor"

func UserCheck(next echo.HandlerFunc) echo.HandlerFunc {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User")

		if strings.EqualFold(val, userAdmin) {
			log.Info().Msg("It's Igor")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
