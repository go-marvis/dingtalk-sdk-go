package dingtalk

import (
	"github.com/go-marvis/dingtalk-sdk-go/core"
	"github.com/go-marvis/dingtalk-sdk-go/service/topapi"
)

type Client struct {
	TopApi *topapi.Service
}

func NewClient(config *core.Config) *Client {
	core.Init(config)

	return &Client{
		topapi.NewService(config),
	}
}
