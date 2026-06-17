package doc

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type docs struct {
	config *core.Config
}

// 创建知识库文档
// https://open.dingtalk.com/document/development/create-team-space-document
func (s *docs) Create(ctx context.Context, req *CreateDocReq, options ...core.RequestOptionFunc) (*CreateDocResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/doc/workspaces/:workspaceId/docs"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &CreateDocResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 删除知识库文档
// https://open.dingtalk.com/document/development/delete-team-space-documents
func (s *docs) Delete(ctx context.Context, req *DeleteDocReq, options ...core.RequestOptionFunc) (*DeleteDocResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodDelete
	apiReq.ApiPath = "/v1.0/doc/workspaces/:workspaceId/docs/:nodeId"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &DeleteDocResp{ApiResp: apiResp}
	return resp, err
}
