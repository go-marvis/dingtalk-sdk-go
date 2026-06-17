package wiki

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type node struct {
	config *core.Config
}

// 获取节点
// https://open.dingtalk.com/document/development/get-knowledge-base-acquisition-node
func (s *node) Get(ctx context.Context, req *GetNodeReq, options ...core.RequestOptionFunc) (*GetNodeResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/v2.0/wiki/nodes/:nodeId"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &GetNodeResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
