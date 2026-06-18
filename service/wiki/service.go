package wiki

import (
	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type Service struct {
	Workspace *workspace
	Node      *node
}

func NewService(config *core.Config) *Service {
	return &Service{
		Workspace: &workspace{config},
		Node:      &node{config},
	}
}
