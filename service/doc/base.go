package doc

import "github.com/go-marvis/dingtalk-sdk-go/core"

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

type request struct {
	apiReq *core.ApiReq
}

type requestBuilder struct {
	apiReq *core.ApiReq
}
