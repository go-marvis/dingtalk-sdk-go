package drive

import "github.com/go-marvis/dingtalk-sdk-go/core"

type Service struct {
	config *core.Config
}

func NewService(config *core.Config) *Service {
	return &Service{config}
}

func (s *Service) GetSpaces() {

}
