package dingtalk

import (
	"github.com/go-marvis/dingtalk-sdk-go/core"
	"github.com/go-marvis/dingtalk-sdk-go/service/topapi"
	wiki "github.com/go-marvis/dingtalk-sdk-go/service/wiki/v2"
)

type Client struct {
	TopApi *topapi.Service
	Wiki   *wiki.Service
}

func NewClient(config *core.Config) *Client {
	core.Init(config)

	return &Client{
		topapi.NewService(config),
		wiki.NewService(config),
	}
}
