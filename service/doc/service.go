package doc

import (
	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type Service struct {
	Docs *docs
}

func NewService(config *core.Config) *Service {
	return &Service{
		Docs: &docs{config},
	}
}
