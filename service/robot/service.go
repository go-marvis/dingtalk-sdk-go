package robot

import "github.com/go-marvis/dingtalk-sdk-go/core"

type Service struct {
	Chat  *chat
	Ding  *ding
	Group *group
	OTO   *oto
}

func NewService(config *core.Config) *Service {
	return &Service{
		&chat{config},
		&ding{config},
		&group{config},
		&oto{config},
	}
}
