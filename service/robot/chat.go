package robot

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type chat struct {
	config *core.Config
}

// 人与人会话中机器人发送普通消息
// https://open.dingtalk.com/document/development/the-robot-sends-ordinary-messages-in-a-person-to-person-conversation
func (s *chat) SendChatMessage(ctx context.Context, req *SendChatMessageReq, options ...core.RequestOptionFunc) (*SendChatMessageResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/robot/privateChatMessages/send"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &SendChatMessageResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
