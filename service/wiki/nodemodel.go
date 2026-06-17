package wiki

import (
	"strconv"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

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
type GetNodeReq request
type GetNodeReqBuilder requestBuilder
type GetNodeResp struct {
	*core.ApiResp
	core.CodeError
	Node *Node `json:"node"`
}

func NewGetNodeReqBuilder() *GetNodeReqBuilder {
	return &GetNodeReqBuilder{core.NewApiReq()}
}

func (b *GetNodeReqBuilder) NodeId(nodeId string) *GetNodeReqBuilder {
	b.apiReq.PathParams.Set("nodeId", nodeId)
	return b
}

func (b *GetNodeReqBuilder) WithStatisticalInfo(withStatisticalInfo bool) *GetNodeReqBuilder {
	b.apiReq.QueryParams.Set("withStatisticalInfo", strconv.FormatBool(withStatisticalInfo))
	return b
}

func (b *GetNodeReqBuilder) WithPermissionRole(withPermissionRole bool) *GetNodeReqBuilder {
	b.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return b
}
func (b *GetNodeReqBuilder) OperatorId(operatorId string) *GetNodeReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *GetNodeReqBuilder) Build() *GetNodeReq {
	return &GetNodeReq{b.apiReq}
}

type ListNodeReq request
type ListNodeReqBuilder requestBuilder
type ListNodeResp struct {
	*core.ApiResp
	core.CodeError
	Nodes     []*Node `json:"nodes"`
	NextToken string  `json:"nextToken"`
}

func NewListNodeReqBuilder() *ListNodeReqBuilder {
	return &ListNodeReqBuilder{core.NewApiReq()}
}

func (b *ListNodeReqBuilder) ParentNodeId(parentNodeId string) *ListNodeReqBuilder {
	b.apiReq.QueryParams.Set("parentNodeId", parentNodeId)
	return b
}

func (b *ListNodeReqBuilder) NextToken(nextToken string) *ListNodeReqBuilder {
	b.apiReq.QueryParams.Set("nextToken", nextToken)
	return b
}

func (b *ListNodeReqBuilder) MaxResults(maxResults int64) *ListNodeReqBuilder {
	b.apiReq.QueryParams.Set("maxResults", strconv.FormatInt(maxResults, 10))
	return b
}

func (b *ListNodeReqBuilder) WithPermissionRole(withPermissionRole bool) *ListNodeReqBuilder {
	b.apiReq.QueryParams.Set("withPermissionRole", strconv.FormatBool(withPermissionRole))
	return b
}

func (b *ListNodeReqBuilder) OperatorId(operatorId string) *ListNodeReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *ListNodeReqBuilder) Build() *ListNodeReq {
	return &ListNodeReq{b.apiReq}
}
