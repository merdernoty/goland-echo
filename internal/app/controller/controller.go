package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}

type Controller struct {
	s Service
}

func New(s Service) *Controller {
	return &Controller{
		s: s,
	}
}

func (e *Controller) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()
	s := fmt.Sprintf("Количество дней: %d", d)
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
