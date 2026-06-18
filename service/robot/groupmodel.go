package robot

import "github.com/go-marvis/dingtalk-sdk-go/core"

type SendGroupMessageReq request
type SendGroupMessageReqBuilder requestBuilder
type SendGroupMessageReqBody struct {
	RobotCode          string `json:"robotCode"`
	MsgKey             string `json:"msgKey"`
	MsgParam           string `json:"msgParam"`
	OpenConversationId string `json:"openConversationId"`
	CoolAppCode        string `json:"coolAppCode,omitempty"`
}

type SendGroupMessageResp struct {
	*core.ApiResp
	core.CodeError
	ProcessQueryKey string `json:"processQueryKey"`
}

func NewSendGroupMessageReqBuilder() *SendGroupMessageReqBuilder {
	return &SendGroupMessageReqBuilder{core.NewApiReq()}
}

func (b *SendGroupMessageReqBuilder) Body(body *SendGroupMessageReqBody) *SendGroupMessageReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *SendGroupMessageReqBuilder) Build() *SendGroupMessageReq {
	return &SendGroupMessageReq{b.apiReq}
}
