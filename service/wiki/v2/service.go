package v2

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

// POST 新建知识库
// https://open.dingtalk.com/document/development/new-knowledge-base
func (s *Service) CreateWorkspace(ctx context.Context, req *CreateWorkspaceReq, options ...core.RequestOptionFunc) (*CreateWorkspaceResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v2.0/wiki/workspaces"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &CreateWorkspaceResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// GET 获取知识库列表
// https://open.dingtalk.com/document/development/get-knowledge-base-list
func (s *Service) GetWorkspaces(ctx context.Context, req *GetWorkspacesReq, options ...core.RequestOptionFunc) (*GetWorkspacesResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/v2.0/wiki/workspaces"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &GetWorkspacesResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// GET 获取知识库
// https://open.dingtalk.com/document/development/obtain-the-knowledge-base
func (s *Service) GetWorkspace(ctx context.Context, req *GetWorkspaceReq, options ...core.RequestOptionFunc) (*GetWorkspaceResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/v2.0/wiki/workspaces/:workspaceId"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &GetWorkspaceResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
