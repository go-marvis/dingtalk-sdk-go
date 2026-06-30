package topapi

import "github.com/go-marvis/dingtalk-sdk-go/core"

// GetUserReq
type GetUserReq request
type GetUserReqBuilder requestBuilder
type GetUserReqBody struct {
	Language string `json:"language"`
	UserId   string `json:"userid"`
}

func NewGetUserReqBuilder() *GetUserReqBuilder {
	return &GetUserReqBuilder{core.NewApiReq()}
}

func (builder *GetUserReqBuilder) Body(body *GetUserReqBody) *GetUserReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *GetUserReqBuilder) Build() *GetUserReq {
	return &GetUserReq{builder.apiReq}
}

type GetUserResp struct {
	*core.ApiResp
	core.CodeError
	Result *GetUserRespResult `json:"result"`
}

type GetUserRespResult struct {
	UserId        string `json:"userid"`
	UnionId       string `json:"unionid"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	StateCode     string `json:"state_code"`
	ManagerUserId string `json:"manager_userid"`
	Mobile        string `json:"mobile"`
	Title         string `json:"title"`
	Email         string `json:"email"`
	WorkPlace     string `json:"work_place"`  // 办公地点
	Remark        string `json:"remark"`      // 备注
	Extension     string `json:"extension"`   // 扩展属性
	Active        bool   `json:"active"`      // 是否激活了钉钉
	RealAuthed    bool   `json:"real_authed"` // 是否完成了实名认证
	Senior        bool   `json:"senior"`      // 是否为企业的高管
	Admin         bool   `json:"admin"`
	Boss          bool   `json:"boss"`
}

// CountUserReq
type CountUserReq request
type CountUserReqBodyBuilder requestBuilder

type CountUserReqBody struct {
	OnlyActive string `json:"only_active,omitempty"`
}

func NewCountUserReqBodyBuilder() *CountUserReqBodyBuilder {
	return &CountUserReqBodyBuilder{core.NewApiReq()}
}

func (b *CountUserReqBodyBuilder) Body(body *CountUserReqBody) *CountUserReqBodyBuilder {
	b.apiReq.Body = body
	return b
}

func (b *CountUserReqBodyBuilder) Build() *CountUserReq {
	return &CountUserReq{b.apiReq}
}

type CountUserResp struct {
	*core.ApiResp
	core.CodeError
	Result *CountUserRespResult `json:"result"`
}

type CountUserRespResult struct {
	Count int64 `json:"count"`
}
