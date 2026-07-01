package robot

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type ding struct {
	config *core.Config
}

// 发送DING消息
// https://open.dingtalk.com/document/development/robot-sends-nail-message
func (s *ding) Send(ctx context.Context, req *SendDingReq, options ...core.RequestOptionFunc) (*SendDingResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/robot/ding/send"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &SendDingResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
