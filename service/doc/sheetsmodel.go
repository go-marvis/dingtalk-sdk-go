package doc

import "github.com/go-marvis/dingtalk-sdk-go/core"

// CreateSheet
type CreateSheetReq request
type CreateSheetReqBuilder requestBuilder

type CreateSheetReqBody struct {
	Name string `json:"name"`
}

type CreateSheetResp struct {
	*core.ApiResp
	core.CodeError
	Visibility string `json:"visibility"`
	Name       string `json:"name"`
	Id         string `json:"id"`
}

func NewCreateSheetReqBuilder() *CreateSheetReqBuilder {
	return &CreateSheetReqBuilder{core.NewApiReq()}
}

func (b *CreateSheetReqBuilder) WorkspaceId(workspaceId string) *CreateSheetReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *CreateSheetReqBuilder) OperatorId(operatorId string) *CreateSheetReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *CreateSheetReqBuilder) Body(body *CreateSheetReqBody) *CreateSheetReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *CreateSheetReqBuilder) Build() *CreateSheetReq {
	return &CreateSheetReq{b.apiReq}
}

type UpdateSheetReq request
type UpdateSheetReqBuilder requestBuilder

type UpdateReqBody struct {
	FrozenRowCount    int64  `json:"frozenRowCount"`
	FrozenColumnCount int64  `json:"frozenColumnCount"`
	Visibility        string `json:"visibility"`
	Name              string `json:"name"`
}

type UpdateSheetResp struct {
	*core.ApiResp
	core.CodeError
	Id string `json:"id"`
}

func NewUpdateSheetReqBuilder() *UpdateSheetReqBuilder {
	return &UpdateSheetReqBuilder{core.NewApiReq()}
}

func (b *UpdateSheetReqBuilder) WorkspaceId(workspaceId string) *UpdateSheetReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *UpdateSheetReqBuilder) SheetId(sheetId string) *UpdateSheetReqBuilder {
	b.apiReq.PathParams.Set("sheetId", sheetId)
	return b
}

func (b *UpdateSheetReqBuilder) OperatorId(operatorId string) *UpdateSheetReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *UpdateSheetReqBuilder) Body(body *UpdateReqBody) *UpdateSheetReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *UpdateSheetReqBuilder) Build() *UpdateSheetReq {
	return &UpdateSheetReq{b.apiReq}
}

type DeleteSheetReq request
type DeleteSheetReqBuilder requestBuilder
type DeleteSheetResp struct {
	*core.ApiResp
	core.CodeError
	Success bool `json:"success"`
}

func NewDeleteSheetReqBuilder() *DeleteSheetReqBuilder {
	return &DeleteSheetReqBuilder{core.NewApiReq()}
}

func (b *DeleteSheetReqBuilder) WorkspaceId(workspaceId string) *DeleteSheetReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *DeleteSheetReqBuilder) SheetId(sheetId string) *DeleteSheetReqBuilder {
	b.apiReq.PathParams.Set("sheetId", sheetId)
	return b
}

func (b *DeleteSheetReqBuilder) OperatorId(operatorId string) *DeleteSheetReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *DeleteSheetReqBuilder) Build() *DeleteSheetReq {
	return &DeleteSheetReq{b.apiReq}
}

type GetSheetReq request
type GetSheetReqBuilder requestBuilder
type GetSheetResp struct {
	*core.ApiResp
	core.CodeError
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Visibility         string `json:"visibility"`
	LastNonEmptyRow    int64  `json:"lastNonEmptyRow"`
	LastNonEmptyColumn int64  `json:"lastNonEmptyColumn"`
	RowCount           int64  `json:"rowCount"`
	ColumnCount        int64  `json:"columnCount"`
}

func NewGetSheetReqBuilder() *GetSheetReqBuilder {
	return &GetSheetReqBuilder{core.NewApiReq()}
}

func (b *GetSheetReqBuilder) WorkspaceId(workspaceId string) *GetSheetReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *GetSheetReqBuilder) SheetId(sheetId string) *GetSheetReqBuilder {
	b.apiReq.PathParams.Set("sheetId", sheetId)
	return b
}

func (b *GetSheetReqBuilder) OperatorId(operatorId string) *GetSheetReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *GetSheetReqBuilder) Build() *GetSheetReq {
	return &GetSheetReq{b.apiReq}
}

type ListSheetReq request
type ListSheetReqBuilder requestBuilder
type ListSheetResp struct {
	*core.ApiResp
	core.CodeError
	Value []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"value"`
}

func NewListSheetReqBuilder() *ListSheetReqBuilder {
	return &ListSheetReqBuilder{core.NewApiReq()}
}

func (b *ListSheetReqBuilder) WorkspaceId(workspaceId string) *ListSheetReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *ListSheetReqBuilder) OperatorId(operatorId string) *ListSheetReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *ListSheetReqBuilder) Build() *ListSheetReq {
	return &ListSheetReq{b.apiReq}
}

type DeleteRowsReq request
type DeleteRowsReqBuilder requestBuilder
type DeleteRowsReqBody struct {
	Row      int64 `json:"row"`
	RowCount int64 `json:"rowCount"`
}

type DeleteRowsResp struct {
	*core.ApiResp
	core.CodeError
	Id string `json:"id"`
}

func NewDeleteRowsReqBuilder() *DeleteRowsReqBuilder {
	return &DeleteRowsReqBuilder{core.NewApiReq()}
}

func (b *DeleteRowsReqBuilder) WorkspaceId(workspaceId string) *DeleteRowsReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *DeleteRowsReqBuilder) SheetId(sheetId string) *DeleteRowsReqBuilder {
	b.apiReq.PathParams.Set("sheetId", sheetId)
	return b
}

func (b *DeleteRowsReqBuilder) OperatorId(operatorId string) *DeleteRowsReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *DeleteRowsReqBuilder) Build() *DeleteRowsReq {
	return &DeleteRowsReq{b.apiReq}
}

type DeleteColumnsReq request
type DeleteColumnsReqBuilder requestBuilder
type DeleteColumnsReqBody struct {
	Column      int64 `json:"column"`
	ColumnCount int64 `json:"columnCount"`
}

type DeleteColumnsResp struct {
	*core.ApiResp
	core.CodeError
	Id string `json:"id"`
}

func NewDeleteColumnsReqBuilder() *DeleteColumnsReqBuilder {
	return &DeleteColumnsReqBuilder{core.NewApiReq()}
}

func (b *DeleteColumnsReqBuilder) WorkspaceId(workspaceId string) *DeleteColumnsReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *DeleteColumnsReqBuilder) SheetId(sheetId string) *DeleteColumnsReqBuilder {
	b.apiReq.PathParams.Set("sheetId", sheetId)
	return b
}

func (b *DeleteColumnsReqBuilder) OperatorId(operatorId string) *DeleteColumnsReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *DeleteColumnsReqBuilder) Build() *DeleteColumnsReq {
	return &DeleteColumnsReq{b.apiReq}
}

type AppendRowsReq request
type AppendRowsReqBuilder requestBuilder
type AppendRowsReqBody struct {
	Values [][]string `json:"values"`
}
type AppendRowsResp struct {
	*core.ApiResp
	core.CodeError
	A1Notation string `json:"a1Notation"`
}

func NewAppendRowsReqBuilder() *AppendRowsReqBuilder {
	return &AppendRowsReqBuilder{core.NewApiReq()}
}

func (b *AppendRowsReqBuilder) WorkspaceId(workspaceId string) *AppendRowsReqBuilder {
	b.apiReq.PathParams.Set("workspaceId", workspaceId)
	return b
}

func (b *AppendRowsReqBuilder) SheetId(sheetId string) *AppendRowsReqBuilder {
	b.apiReq.PathParams.Set("sheetId", sheetId)
	return b
}

func (b *AppendRowsReqBuilder) OperatorId(operatorId string) *AppendRowsReqBuilder {
	b.apiReq.QueryParams.Set("operatorId", operatorId)
	return b
}

func (b *AppendRowsReqBuilder) Body(body AppendRowsReqBody) *AppendRowsReqBuilder {
	b.apiReq.Body = body
	return b
}

func (b *AppendRowsReqBuilder) Build() *AppendRowsReq {
	return &AppendRowsReq{b.apiReq}
}
