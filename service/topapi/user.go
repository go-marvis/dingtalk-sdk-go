package topapi

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type user struct {
	config *core.Config
}

// 查询用户详情
// https://open.dingtalk.com/document/development/query-user-details
func (s *user) Get(ctx context.Context, req *GetUserReq, options ...core.RequestOptionFunc) (*GetUserResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = s.config.OapiUrl + "/topapi/v2/user/get"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &GetUserResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 获取员工人数
// https://open.dingtalk.com/document/development/user-management-acquires-number-employees
func (s *user) Count(ctx context.Context, req *CountUserReq, options ...core.RequestOptionFunc) (*CountUserResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = s.config.OapiUrl + "/topapi/user/count"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &CountUserResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
