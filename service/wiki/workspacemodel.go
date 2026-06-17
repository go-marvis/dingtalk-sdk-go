package wiki

import (
	"strconv"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

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
type CreateWorkspaceReq request
type CreateWorkspaceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewCreateWorkspaceReqBuilder() *CreateWorkspaceReqBuilder {
	return &CreateWorkspaceReqBuilder{core.NewApiReq()}
}

func (b *CreateWorkspaceReqBuilder) Body(body *CreateWorkspaceReqBody) *CreateWorkspaceReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *CreateWorkspaceReqBuilder) Build() *CreateWorkspaceReq {
	return &CreateWorkspaceReq{b.apiReq}
}

type CreateWorkspaceReqBody struct {
	Name   string `json:"name"`
	Option struct {
		Description string `json:"description,omitempty"`
		TeamId      string `json:"teamId,omitempty"`
	} `json:"option"`
}

type CreateWorkspaceResp struct {
	*core.ApiResp
	core.CodeError
	Workspace *Workspace `json:"workspace"`
}

// GetWorkspaceReq
type GetWorkspaceReq request

type GetWorkspaceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetWorkspaceReqBuilder() *GetWorkspaceReqBuilder {
	return &GetWorkspaceReqBuilder{core.NewApiReq()}
}

func (b *GetWorkspaceReqBuilder) WorkspaceId(workspaceId string) *GetWorkspaceReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *GetWorkspaceReqBuilder) WithPermissionRole(withPermissionRole bool) *GetWorkspaceReqBuilder {
	b.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return b
}

func (b *GetWorkspaceReqBuilder) OperatorId(operatorId string) *GetWorkspaceReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *GetWorkspaceReqBuilder) Build() *GetWorkspaceReq {
	return &GetWorkspaceReq{b.apiReq}
}

type GetWorkspaceResp struct {
	*core.ApiResp
	core.CodeError
	Workspace *Workspace `json:"workspace"`
}

// ListWorkspaceReq
type ListWorkspaceReq request
type ListWorkspaceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetWorkspacesReqBuilder() *ListWorkspaceReqBuilder {
	return &ListWorkspaceReqBuilder{core.NewApiReq()}
}

func (b *ListWorkspaceReqBuilder) NextToken(nextToken string) *ListWorkspaceReqBuilder {
	b.apiReq.QueryParams.Set("nextToken", nextToken)
	return b
}

func (b *ListWorkspaceReqBuilder) MaxResults(maxResults int64) *ListWorkspaceReqBuilder {
	b.apiReq.QueryParams.Set("maxResults", strconv.FormatInt(maxResults, 10))
	return b
}

func (b *ListWorkspaceReqBuilder) OrderBy(orderBy string) *ListWorkspaceReqBuilder {
	b.apiReq.QueryParams.Set("orderBy", orderBy)
	return b
}

func (b *ListWorkspaceReqBuilder) WithPermissionRole(withPermissionRole bool) *ListWorkspaceReqBuilder {
	b.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return b
}

func (b *ListWorkspaceReqBuilder) OperatorId(operatorId string) *ListWorkspaceReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *ListWorkspaceReqBuilder) Build() *ListWorkspaceReq {
	return &ListWorkspaceReq{b.apiReq}
}

type ListWorkspaceResp struct {
	*core.ApiResp
	core.CodeError
	Workspaces []*Workspace `json:"workspaces"`
	NextToken  string       `json:"next_token"`
}

// MineWorkspaces
type MineWorkspaceReq request
type MineWorkspaceReqBuilder requestBuilder

func NewMineWorkspacesReqBuilder() *MineWorkspaceReqBuilder {
	return &MineWorkspaceReqBuilder{core.NewApiReq()}
}

func (b *MineWorkspaceReqBuilder) OperatorId(operatorId string) *MineWorkspaceReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *MineWorkspaceReqBuilder) Build() *MineWorkspaceReq {
	return &MineWorkspaceReq{b.apiReq}
}

type MineWorkspaceResp struct {
	*core.ApiResp
	core.CodeError
	Workspace *Workspace `json:"workspace"`
}
