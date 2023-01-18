package service

import (
	"math"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) CalculateDays(year int) int {
	date := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	daysToYear := int(math.Abs(time.Until(date).Hours() / 24))

	return daysToYear
}
