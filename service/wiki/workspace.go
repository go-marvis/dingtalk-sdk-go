package wiki

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type workspace struct {
	config *core.Config
}

// 新建知识库
// https://open.dingtalk.com/document/development/new-knowledge-base
func (s *workspace) Create(ctx context.Context, req *CreateWorkspaceReq, options ...core.RequestOptionFunc) (*CreateWorkspaceResp, error) {
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

// 获取知识库
// https://open.dingtalk.com/document/development/obtain-the-knowledge-base
func (s *workspace) Get(ctx context.Context, req *GetWorkspaceReq, options ...core.RequestOptionFunc) (*GetWorkspaceResp, error) {
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

// 获取知识库列表
// https://open.dingtalk.com/document/development/get-knowledge-base-list
func (s *workspace) List(ctx context.Context, req *ListWorkspaceReq, options ...core.RequestOptionFunc) (*ListWorkspaceResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/v2.0/wiki/workspaces"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &ListWorkspaceResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 获取我的文档知识库信息
// https://open.dingtalk.com/document/development/get-my-documents
func (s *workspace) Mine(ctx context.Context, req *MineWorkspaceReq, options ...core.RequestOptionFunc) (*MineWorkspaceResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/v2.0/wiki/mineWorkspaces"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &MineWorkspaceResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
