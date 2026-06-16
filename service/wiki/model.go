package wiki

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

// ListWorkspaceReq
type ListWorkspaceReq Request
type ListWorkspaceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetWorkspacesReqBuilder() *ListWorkspaceReqBuilder {
	return &ListWorkspaceReqBuilder{core.NewApiReq()}
}

func (builder *ListWorkspaceReqBuilder) NextToken(nextToken string) *ListWorkspaceReqBuilder {
	builder.apiReq.QueryParams.Set("nextToken", nextToken)
	return builder
}

func (builder *ListWorkspaceReqBuilder) MaxResults(maxResults int64) *ListWorkspaceReqBuilder {
	builder.apiReq.QueryParams.Set("maxResults", strconv.FormatInt(maxResults, 10))
	return builder
}

func (builder *ListWorkspaceReqBuilder) OrderBy(orderBy string) *ListWorkspaceReqBuilder {
	builder.apiReq.QueryParams.Set("orderBy", orderBy)
	return builder
}

func (builder *ListWorkspaceReqBuilder) WithPermissionRole(withPermissionRole bool) *ListWorkspaceReqBuilder {
	builder.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return builder
}

func (builder *ListWorkspaceReqBuilder) OperatorId(operatorId string) *ListWorkspaceReqBuilder {
	builder.apiReq.QueryParams.Set("operatorId", operatorId)
	return builder
}

func (builder *ListWorkspaceReqBuilder) Build() *ListWorkspaceReq {
	return &ListWorkspaceReq{builder.apiReq}
}

type ListWorkspaceResp struct {
	*core.ApiResp
	core.CodeError
	Workspaces []*Workspace `json:"workspaces"`
	NextToken  string       `json:"next_token"`
}

// MineWorkspaces
type MineWorkspaceReq Request

type MineWorkspaceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewMineWorkspacesReqBuilder() *MineWorkspaceReqBuilder {
	return &MineWorkspaceReqBuilder{core.NewApiReq()}
}

func (builder *MineWorkspaceReqBuilder) OperatorId(operatorId string) *MineWorkspaceReqBuilder {
	builder.apiReq.QueryParams.Set("operatorId", operatorId)
	return builder
}

func (builder *MineWorkspaceReqBuilder) Build() *MineWorkspaceReq {
	return &MineWorkspaceReq{builder.apiReq}
}

type MineWorkspaceResp struct {
	*core.ApiResp
	core.CodeError
	Workspace *Workspace `json:"workspace"`
}

type StatisticalInfo struct {
	WordCount int64 `json:"wordCount"`
}

type Node struct {
	NodeId          string           `json:"nodeId"`
	WorkspaceeId    string           `json:"workspaceId"`
	Name            string           `json:"name"`
	Size            int64            `json:"size"`
	Type            string           `json:"type"`
	Category        string           `json:"category"`
	Extension       string           `json:"extension"`
	Url             string           `json:"url"`
	CreatorId       string           `json:"creatorId"`
	ModifierId      string           `json:"modifierId"`
	CreateTime      string           `json:"createTime"`
	ModifiedTime    string           `json:"modifiedTime"`
	HasChildren     bool             `json:"hasChildren"`
	StatisticalInfo *StatisticalInfo `json:"statisticalInfo"`
	PermissionRole  string           `json:"permissionRole"`
}

// GetNodeReq
type GetNodeReq Request
type GetNodeReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetNodeReqBuilder() *GetNodeReqBuilder {
	return &GetNodeReqBuilder{core.NewApiReq()}
}

func (builder *GetNodeReqBuilder) NodeId(nodeId string) *GetNodeReqBuilder {
	builder.apiReq.PathParams.Set("nodeId", nodeId)
	return builder
}

func (builder *GetNodeReqBuilder) WithStatisticalInfo(withStatisticalInfo bool) *GetNodeReqBuilder {
	builder.apiReq.QueryParams.Set("withStatisticalInfo", strconv.FormatBool(withStatisticalInfo))
	return builder
}

func (builder *GetNodeReqBuilder) WithPermissionRole(withPermissionRole bool) *GetNodeReqBuilder {
	builder.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return builder
}
func (builder *GetNodeReqBuilder) OperatorId(operatorId string) *GetNodeReqBuilder {
	builder.apiReq.QueryParams.Set("operatorId", operatorId)
	return builder
}

func (builder *GetNodeReqBuilder) Build() *GetNodeReq {
	return &GetNodeReq{builder.apiReq}
}

type GetNodeResp struct {
	*core.ApiResp
	core.CodeError
	Node *Node `json:"node"`
}

type ListNodeReq Request
type ListNodeReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListNodeReqBuilder() *ListNodeReqBuilder {
	return &ListNodeReqBuilder{core.NewApiReq()}
}

func (builder *ListNodeReqBuilder) ParentNodeId(parentNodeId string) *ListNodeReqBuilder {
	builder.apiReq.QueryParams.Set("parentNodeId", parentNodeId)
	return builder
}

func (builder *ListNodeReqBuilder) NextToken(nextToken string) *ListNodeReqBuilder {
	builder.apiReq.QueryParams.Set("nextToken", nextToken)
	return builder
}

func (builder *ListNodeReqBuilder) MaxResults(maxResults int64) *ListNodeReqBuilder {
	builder.apiReq.QueryParams.Set("maxResults", strconv.FormatInt(maxResults, 10))
	return builder
}

func (builder *ListNodeReqBuilder) WithPermissionRole(withPermissionRole bool) *ListNodeReqBuilder {
	builder.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return builder
}

func (builder *ListNodeReqBuilder) OperatorId(operatorId string) *ListNodeReqBuilder {
	builder.apiReq.QueryParams.Set("operatorId", operatorId)
	return builder
}

func (builder *ListNodeReqBuilder) Build() *ListNodeReq {
	return &ListNodeReq{builder.apiReq}
}

type ListNodeResp struct {
	*core.ApiResp
	core.CodeError
	Nodes     []*Node `json:"nodes"`
	NextToken string  `json:"nextToken"`
}
