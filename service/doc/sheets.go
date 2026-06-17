package doc

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type sheets struct {
	config *core.Config
}

// 创建工作表
// https://open.dingtalk.com/document/development/create-a-worksheet
func (s *sheets) Create(ctx context.Context, req *CreateSheetReq, options ...core.RequestOptionFunc) (*CreateSheetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/doc/workbooks/:workbookId/sheets"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &CreateSheetResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 更新工作表
// https://open.dingtalk.com/document/development/create-a-worksheet
func (s *sheets) Update(ctx context.Context, req *UpdateSheetReq, options ...core.RequestOptionFunc) (*UpdateSheetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/doc/workbooks/:workbookId/sheets/:sheetId"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &UpdateSheetResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 删除工作表
// https://open.dingtalk.com/document/development/delete-classic-workbooks
func (s *sheets) Delete(ctx context.Context, req *DeleteSheetReq, options ...core.RequestOptionFunc) (*DeleteSheetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/doc/workbooks/:workbookId/sheets/:sheetId"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &DeleteSheetResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 获取工作表
// https://open.dingtalk.com/document/development/obtain-worksheet-properties
func (s *sheets) Get(ctx context.Context, req *GetSheetReq, options ...core.RequestOptionFunc) (*GetSheetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/v1.0/doc/workbooks/:workbookId/sheets/:sheetId"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &GetSheetResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 获取所有工作表
// https://open.dingtalk.com/document/development/obtain-all-worksheets
func (s *sheets) List(ctx context.Context, req *ListSheetReq, options ...core.RequestOptionFunc) (*ListSheetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodGet
	apiReq.ApiPath = "/v1.0/doc/workbooks/:workbookId/sheets"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &ListSheetResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 删除行
// https://open.dingtalk.com/document/development/delete-row
func (s *sheets) DeleteRows(ctx context.Context, req *DeleteRowsReq, options ...core.RequestOptionFunc) (*DeleteRowsResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/doc/workbooks/:workbookId/sheets/:sheetId/deleteRows"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &DeleteRowsResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 删除列
// https://open.dingtalk.com/document/development/delete-column
func (s *sheets) DeleteColumns(ctx context.Context, req *DeleteColumnsReq, options ...core.RequestOptionFunc) (*DeleteColumnsResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/doc/workbooks/:workbookId/sheets/:sheetId/deleteColumns"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &DeleteColumnsResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// 工作表中追加行
// https://open.dingtalk.com/document/development/append-line
func (s *sheets) AppendRows(ctx context.Context, req *AppendRowsReq, options ...core.RequestOptionFunc) (*AppendRowsResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/doc/workbooks/:workbookId/sheets/:sheetId/appendRows"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &AppendRowsResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
