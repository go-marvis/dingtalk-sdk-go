package robot

import (
	"encoding/json"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type BatchSendOTOReq request
type BatchSendOTOReqBuilder requestBuilder
type BatchSendOTOReqBody struct {
	RobotCode string   `json:"robotCode"`
	UserIds   []string `json:"userIds"`
	MsgKey    string   `json:"msgKey"`
	MsgParam  string   `json:"msgParam"`
}

func (b *BatchSendOTOReqBody) SetParam(param any) {
	data, _ := json.Marshal(param)
	b.MsgParam = string(data)
}

type BatchSendOTOResp struct {
	*core.ApiResp
	core.CodeError
	ProcessQueryKey           string   `json:"processQueryKey"`
	InvalidStaffIdList        []string `json:"invalidStaffIdList"`
	FlowControlledStaffIdList []string `json:"flowControlledStaffIdList"`
}

func NewBatchSendOTOReqBuilder() *BatchSendOTOReqBuilder {
	return &BatchSendOTOReqBuilder{core.NewApiReq()}
}

func (b *BatchSendOTOReqBuilder) Body(body *BatchSendOTOReqBody) *BatchSendOTOReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *BatchSendOTOReqBuilder) Build() *BatchSendOTOReq {
	return &BatchSendOTOReq{b.apiReq}
}
