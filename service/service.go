package service

import (
	"github.com/go-kit/log"
)

type Service interface {
}

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}
