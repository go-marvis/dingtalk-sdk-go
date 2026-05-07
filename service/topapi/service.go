package topapi

import (
	"github.com/go-marvis/dingtalk-sdk-go/core"
	"github.com/go-marvis/dingtalk-sdk-go/service/topapi/user"
)

type Service struct {
	User *user.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		User: user.NewService(config),
	}
}
