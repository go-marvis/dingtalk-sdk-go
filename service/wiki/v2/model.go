package v2

import (
	"strconv"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type Request struct {
	apiReq *core.ApiReq
}

type Icon struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Workspace struct {
	WorkspaceId    string `json:"workspaceId"`
	CorpId         string `json:"corpId"`
	TeamId         string `json:"teamId"`
	RootNodeId     string `json:"rootNodeId"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Icon           Icon   `json:"icon"`
	Cover          string `json:"cover"`
	CreatorId      string `json:"creatorId"`
	ModifierId     string `json:"modifierId"`
	CreateTime     string `json:"createTime"`
	ModifiedTime   string `json:"modifiedTime"`
	PermissionRole string `json:"permissionRole"`
}

// CreateWorkspaceReq
type CreateWorkspaceReq Request
type CreateWorkspaceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewCreateWorkspaceReqBuilder() *CreateWorkspaceReqBuilder {
	return &CreateWorkspaceReqBuilder{core.NewApiReq()}
}

func (builder *CreateWorkspaceReqBuilder) Body(body *CreateWorkspaceReqBody) *CreateWorkspaceReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *CreateWorkspaceReqBuilder) Build() *CreateWorkspaceReq {
	return &CreateWorkspaceReq{builder.apiReq}
}

type CreateWorkspaceReqBody struct {
	Name   string `json:"name"`
	Option struct {
		Description string `json:"description"`
		TeamId      string `json:"teamId"`
	} `json:"option"`
}

type CreateWorkspaceResp struct {
	*core.ApiResp
	core.CodeError
	Workspace *Workspace `json:"workspace"`
}

// GetWorkspaceReq
type GetWorkspaceReq Request

type GetWorkspaceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetWorkspaceReqBuilder() *GetWorkspaceReqBuilder {
	return &GetWorkspaceReqBuilder{core.NewApiReq()}
}

func (builder *GetWorkspaceReqBuilder) WorkspaceId(workspaceId string) *GetWorkspaceReqBuilder {
	builder.apiReq.PathParams.Set("workspaceId", workspaceId)
	return builder
}

func (builder *GetWorkspaceReqBuilder) WithPermissionRole(withPermissionRole bool) *GetWorkspaceReqBuilder {
	builder.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return builder
}

func (builder *GetWorkspaceReqBuilder) OperatorId(operatorId string) *GetWorkspaceReqBuilder {
	builder.apiReq.QueryParams.Set("operatorId", operatorId)
	return builder
}

func (builder *GetWorkspaceReqBuilder) Build() *GetWorkspaceReq {
	return &GetWorkspaceReq{builder.apiReq}
}

type GetWorkspaceResp struct {
	*core.ApiResp
	core.CodeError
	Workspace *Workspace `json:"workspace"`
}

// GetWorkspacesReq
type GetWorkspacesReq Request
type GetWorkspacesReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetWorkspacesReqBuilder() *GetWorkspacesReqBuilder {
	return &GetWorkspacesReqBuilder{core.NewApiReq()}
}

func (builder *GetWorkspacesReqBuilder) NextToken(nextToken string) *GetWorkspacesReqBuilder {
	builder.apiReq.QueryParams.Set("nextToken", nextToken)
	return builder
}

func (builder *GetWorkspacesReqBuilder) MaxResults(maxResults int64) *GetWorkspacesReqBuilder {
	builder.apiReq.QueryParams.Set("maxResults", strconv.FormatInt(maxResults, 10))
	return builder
}

func (builder *GetWorkspacesReqBuilder) OrderBy(orderBy string) *GetWorkspacesReqBuilder {
	builder.apiReq.QueryParams.Set("orderBy", orderBy)
	return builder
}

func (builder *GetWorkspacesReqBuilder) WithPermissionRole(withPermissionRole bool) *GetWorkspacesReqBuilder {
	builder.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return builder
}

func (builder *GetWorkspacesReqBuilder) OperatorId(operatorId string) *GetWorkspacesReqBuilder {
	builder.apiReq.QueryParams.Set("operatorId", operatorId)
	return builder
}

func (builder *GetWorkspacesReqBuilder) Build() *GetWorkspacesReq {
	return &GetWorkspacesReq{builder.apiReq}
}

type GetWorkspacesResp struct {
	*core.ApiResp
	core.CodeError
	Workspaces []*Workspace `json:"workspaces"`
	NextToken  string       `json:"next_token"`
}

// MineWorkspaces
type MineWorkspacesReq Request

type MineWorkspacesReqBuilder struct {
	apiReq *core.ApiReq
}

func NewMineWorkspacesReqBuilder() *MineWorkspacesReqBuilder {
	return &MineWorkspacesReqBuilder{core.NewApiReq()}
}

func (builder *MineWorkspacesReqBuilder) OperatorId(operatorId string) *MineWorkspacesReqBuilder {
	builder.apiReq.QueryParams.Set("operatorId", operatorId)
	return builder
}

func (builder *MineWorkspacesReqBuilder) Build() *MineWorkspacesReq {
	return &MineWorkspacesReq{builder.apiReq}
}

type MineWorkspacesResp struct {
	*core.ApiResp
	core.CodeError
	Workspace *Workspace `json:"workspace"`
}
