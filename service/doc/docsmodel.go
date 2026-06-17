package doc

import "github.com/go-marvis/dingtalk-sdk-go/core"

// CreateDoc
type CreateDocReq request
type CreateDocReqBuilder requestBuilder
type CreateDocReqBody struct {
	Name         string       `json:"name"`
	DocType      DocType      `json:"docType"`
	OperatorId   string       `json:"operatorId"`
	ParentNodeId string       `json:"parentNodeId,omitempty"`
	TemplateId   string       `json:"templateId,omitempty"`
	TemplateType TemplateType `json:"templateType,omitempty"`
}

type CreateDocResp struct {
	*core.ApiResp
	core.CodeError
	WorkspaceId string `json:"workspaceId"`
	NodeId      string `json:"nodeId"`
	DocKey      string `json:"docKey"`
	Url         string `json:"url"`
	DentryUuid  string `json:"dentryUuid"`
}

func NewCreateDocReqBuilder() *CreateDocReqBuilder {
	return &CreateDocReqBuilder{core.NewApiReq()}
}

func (b *CreateDocReqBuilder) WorkspaceId(workspaceId string) *CreateDocReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *CreateDocReqBuilder) Body(body *CreateDocReqBody) *CreateDocReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *CreateDocReqBuilder) Build() *CreateDocReq {
	return &CreateDocReq{b.apiReq}
}

// DeleteDoc
type DeleteDocReq request
type DeleteDocReqBuilder requestBuilder
type DeleteDocResp struct {
	*core.ApiResp
	core.CodeError
}

func NewDeleteDocReqBuilder() *DeleteDocReqBuilder {
	return &DeleteDocReqBuilder{core.NewApiReq()}
}

func (b *DeleteDocReqBuilder) WorkspaceId(workspaceId string) *DeleteDocReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *DeleteDocReqBuilder) Operatorid(operatorId string) *DeleteDocReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *DeleteDocReqBuilder) NodeId(nodeId string) *DeleteDocReqBuilder {
	b.apiReq.PathParams.Set("nodeId", nodeId)
	return b
}

func (b *DeleteDocReqBuilder) Build() *DeleteDocReq {
	return &DeleteDocReq{b.apiReq}
}
