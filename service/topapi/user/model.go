package user

import "github.com/go-marvis/dingtalk-sdk-go/core"

type GetReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetReqBuilder() *GetReqBuilder {
	return &GetReqBuilder{core.NewApiReq()}
}

func (builder *GetReqBuilder) Body(body *GetReqBody) *GetReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *GetReqBuilder) Build() *GetReq {
	return &GetReq{builder.apiReq}
}

type GetReq struct {
	apiReq *core.ApiReq
}

type GetReqBody struct {
	Language string `json:"language"`
	UserId   string `json:"userid"`
}

type GetResp struct {
	*core.ApiResp
	core.CodeError
	Result *GetRespResult `json:"result"`
}

type GetRespResult struct {
	UserId string `json:"userid"`
	Name   string `json:"name"`
}
