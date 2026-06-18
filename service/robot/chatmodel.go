package robot

import "github.com/go-marvis/dingtalk-sdk-go/core"

type SendChatMessageReq request
type SendChatMessageReqBuilder requestBuilder
type SendChatMessageReqBody struct {
	RobotCode          string `json:"robotCode"`
	MsgKey             string `json:"msgKey"`
	MsgParam           string `json:"msgParam"`
	OpenConversationId string `json:"openConversationId"`
	CoolAppCode        string `json:"coolAppCode"`
}

type SendChatMessageResp struct {
	*core.ApiResp
	core.CodeError
	ProcessQueryKey string `json:"processQueryKey"`
}

func NewSendChatMessageReqBuilder() *SendChatMessageReqBuilder {
	return &SendChatMessageReqBuilder{core.NewApiReq()}
}

func (b *SendChatMessageReqBuilder) Body(body *SendChatMessageReqBody) *SendChatMessageReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *SendChatMessageReqBuilder) Build() *SendChatMessageReq {
	return &SendChatMessageReq{b.apiReq}
}
