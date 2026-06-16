package core

import (
	"net/http"
	"net/url"
)

type ApiReq struct {
	HttpMethod                string
	ApiPath                   string
	Body                      any
	QueryParams               QueryParams
	PathParams                PathParams
	SupportedAccessTokenTypes []AccessTokenType
}

func NewApiReq() *ApiReq {
	return &ApiReq{
		PathParams:  PathParams{},
		QueryParams: QueryParams{},
	}
}

type PathParams map[string]string

func (p PathParams) Get(key string) string {
	return p[key]
}
func (p PathParams) Set(key, val string) {
	p[key] = val
}

type QueryParams = url.Values

type RequestOption struct {
	AppAccessToken   string // 企业内部应用
	CorpAccessToken  string // 第三方应用授权企业
	UserAccessToken  string // 第三方个人应用
	SuiteAccessToken string // 第三方企业应用

	Header http.Header
}

type RequestOptionFunc func(option *RequestOption)

func WithHeader(header http.Header) RequestOptionFunc {
	return func(option *RequestOption) {
		option.Header = header
	}
}

func WithAppAccessToken(accessToken string) RequestOptionFunc {
	return func(option *RequestOption) {
		option.AppAccessToken = accessToken
	}
}
