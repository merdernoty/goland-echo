package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	d := time.Date(2025, time.March, 10, 0, 0, 0, 0, time.UTC) // My birthday
	dur := time.Until(d)
	return int64(dur.Hours()) / 24
}
