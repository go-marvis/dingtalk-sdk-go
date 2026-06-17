package robot

import "github.com/go-marvis/dingtalk-sdk-go/core"

type request struct {
	apiReq *core.ApiReq
}

type requestBuilder struct {
	apiReq *core.ApiReq
}

type Service struct {
	Ding *ding
}

func NewService(config *core.Config) *Service {
	return &Service{
		&ding{config},
	}
}
