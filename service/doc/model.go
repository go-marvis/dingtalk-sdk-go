package doc

import "github.com/go-marvis/dingtalk-sdk-go/core"

type Request struct {
	apiReq *core.ApiReq
}

type DocType string

const (
	DocTypeDoc      DocType = "DOC"
	DocTypeWorkbook DocType = "WORKBOOK"
	DocTypeMind     DocType = "MIND"
	DocTypeFolder   DocType = "FOLDER"
)

type TemplateType string

const (
	TemplateTypePublic TemplateType = "public_template"
	TemplateTypeTeam   TemplateType = "team_template"
	TemplateTypeUser   TemplateType = "user_template"
)

type CreateDocReqBody struct {
	Name         string       `json:"name"`
	DocType      DocType      `json:"docType"`
	OperatorId   string       `json:"operatorId"`
	ParentNodeId string       `json:"parentNodeId"`
	TemplateId   string       `json:"templateId"`
	TemplateType TemplateType `json:"templateType"`
}

// CreateDocReq
type CreateDocReq Request
type CreateDocReqBuilder struct {
	apiReq *core.ApiReq
}

func NewCreateDocReqBuilder() *CreateDocReqBuilder {
	return &CreateDocReqBuilder{core.NewApiReq()}
}

func (builder *CreateDocReqBuilder) WorkspaceId(workspaceId string) *CreateDocReqBuilder {
	builder.apiReq.PathParams.Set("workspaceId", workspaceId)
	return builder
}

func (builder *CreateDocReqBuilder) Body(body *CreateDocReqBody) *CreateDocReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *CreateDocReqBuilder) Build() *CreateDocReq {
	return &CreateDocReq{builder.apiReq}
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

// DeleteDocReq
type DeleteDocReq Request
type DeleteDocReqBuilder struct {
	apiReq *core.ApiReq
}

func NewDeleteDocReqBuilder() *DeleteDocReqBuilder {
	return &DeleteDocReqBuilder{core.NewApiReq()}
}

func (builder *DeleteDocReqBuilder) WorkspaceId(workspaceId string) *DeleteDocReqBuilder {
	builder.apiReq.PathParams.Set("workspaceId", workspaceId)
	return builder
}
func (builder *DeleteDocReqBuilder) NodeId(nodeId string) *DeleteDocReqBuilder {
	builder.apiReq.PathParams.Set("nodeId", nodeId)
	return builder
}

func (builder *DeleteDocReqBuilder) Operatorid(operatorId string) *DeleteDocReqBuilder {
	builder.apiReq.QueryParams.Set("operatorId", operatorId)
	return builder
}

func (builder *DeleteDocReqBuilder) Build() *DeleteDocReq {
	return &DeleteDocReq{builder.apiReq}
}

type DeleteDocResp struct {
	*core.ApiResp
	core.CodeError
}
