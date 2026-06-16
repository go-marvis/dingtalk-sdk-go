package user

import "github.com/go-marvis/dingtalk-sdk-go/core"

type Request struct {
	apiReq *core.ApiReq
}

// GetReq
type GetReq Request

type GetReqBody struct {
	Language string `json:"language"`
	UserId   string `json:"userid"`
}

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

type GetResp struct {
	*core.ApiResp
	core.CodeError
	Result *GetRespResult `json:"result"`
}

type GetRespResult struct {
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

// CountReq
type CountReq Request

type CountReqBody struct {
	OnlyActive string `json:"only_active,omitempty"`
}

type CountReqBodyBuilder struct {
	apiReq *core.ApiReq
}

func NewCountReqBodyBuilder() *CountReqBodyBuilder {
	return &CountReqBodyBuilder{core.NewApiReq()}
}

func (b *CountReqBodyBuilder) Body(body *CountReqBody) *CountReqBodyBuilder {
	b.apiReq.Body = body
	return b
}

func (b *CountReqBodyBuilder) Build() *CountReq {
	return &CountReq{b.apiReq}
}

type CountResp struct {
	*core.ApiResp
	core.CodeError
	Result *CountRespResult `json:"result"`
}

type CountRespResult struct {
	Count int64 `json:"count"`
}
