package user

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type Service struct {
	config *core.Config
}

func NewService(config *core.Config) *Service {
	return &Service{config}
}

// Get 查询用户详情
// https://open.dingtalk.com/document/development/query-user-details
func (s *Service) Get(ctx context.Context, req *GetReq, options ...core.ReqOptionFunc) (*GetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/topapi/v2/user/get"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &GetResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
