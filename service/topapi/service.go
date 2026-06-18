package topapi

import (
	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type Service struct {
	User *user
}

func NewService(config *core.Config) *Service {
	return &Service{
		User: &user{config},
	}
}
