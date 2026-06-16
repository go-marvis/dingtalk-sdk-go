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
func (s *Service) Get(ctx context.Context, req *GetReq, options ...core.RequestOptionFunc) (*GetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = core.OldUrl + "/topapi/v2/user/get"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &GetResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// Count 获取员工人数
func (s *Service) Count(ctx context.Context, req *CountReq, options ...core.RequestOptionFunc) (*CountResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = core.OldUrl + "/topapi/user/count"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &CountResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
