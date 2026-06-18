package robot

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type group struct {
	config *core.Config
}

// 机器人发送群聊消息
// https://open.dingtalk.com/document/development/the-robot-sends-a-group-message
func (s *group) SendGroupMessage(ctx context.Context, req *SendGroupMessageReq, options ...core.RequestOptionFunc) (*SendGroupMessageResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/robot/groupMessages/send"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &SendGroupMessageResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
