package robot

import "github.com/go-marvis/dingtalk-sdk-go/core"

type SendDingReq request
type SendDingReqBuilder requestBuilder
type SendDingReqBody struct {
	RobotCode          string   `json:"robotCode"`
	RemindType         int64    `json:"remindType"`
	ReceiverUserIdList []string `json:"receiverUserIdList"`
	Content            string   `json:"content"`
	CallVoice          string   `json:"callVoice,omitempty"`
}
type SendDingResp struct {
	*core.ApiResp
	core.CodeError
	OpenDingId string `json:"openDingId"`
}

func NewSendDingReqBuilder() *SendDingReqBuilder {
	return &SendDingReqBuilder{core.NewApiReq()}
}

func (b *SendDingReqBuilder) Body(body *SendDingReqBody) *SendDingReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *SendDingReqBuilder) Build() *SendDingReq {
	return &SendDingReq{b.apiReq}
}
