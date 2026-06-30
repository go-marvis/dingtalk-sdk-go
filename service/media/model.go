package media

import "github.com/go-marvis/dingtalk-sdk-go/core"

type request struct {
	apiReq *core.ApiReq
}

type requestBuilder struct {
	apiReq *core.ApiReq
}

type UploadReq request
type UploadReqBuilder requestBuilder

type UploadResp struct {
	*core.ApiResp
	core.CodeError
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

func NewUploadReqBuilder() *UploadReqBuilder {
	return &UploadReqBuilder{core.NewApiReq()}
}

func (builder *UploadReqBuilder) Body(body *core.Formdata) *UploadReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *UploadReqBuilder) Build() *UploadReq {
	return &UploadReq{builder.apiReq}
}
