package dingtalk

import (
	"github.com/go-marvis/dingtalk-sdk-go/core"
	"github.com/go-marvis/dingtalk-sdk-go/service/doc"
	"github.com/go-marvis/dingtalk-sdk-go/service/robot"
	"github.com/go-marvis/dingtalk-sdk-go/service/topapi"
	"github.com/go-marvis/dingtalk-sdk-go/service/wiki"
)

type Client struct {
	Doc    *doc.Service
	Robot  *robot.Service
	TopApi *topapi.Service
	Wiki   *wiki.Service
}

func NewClient(config *core.Config) *Client {
	core.Init(config)

	return &Client{
		doc.NewService(config),
		robot.NewService(config),
		topapi.NewService(config),
		wiki.NewService(config),
	}
}
