package drive

import "github.com/go-marvis/dingtalk-sdk-go/core"

type GetSpacesReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetReqBuilder() *GetSpacesReqBuilder {
	return &GetSpacesReqBuilder{core.NewApiReq()}
}

func (b *GetSpacesReqBuilder) Build() *GetSpaceReq {
	return &GetSpaceReq{b.apiReq}
}

type GetSpaceReq struct {
	apiReq *core.ApiReq
}

type Space struct {
}

type GetSpacesResult struct {
	Language string `json:"language"`
	UserId   string `json:"userid"`
}

type GetResp struct {
	*core.ApiResp
	core.CodeError
	Result *GetSpacesResult `json:"result"`
}
