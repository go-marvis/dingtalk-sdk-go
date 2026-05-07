package core

import (
	"net/http"
	"net/url"
)

type ApiReq struct {
	HttpMethod      string
	ApiPath         string
	Body            any
	QueryParams     url.Values
	AccessTokenType AccessTokenType
}

func NewApiReq() *ApiReq {
	return &ApiReq{
		QueryParams: url.Values{},
	}
}

type ReqOption struct {
	AccessToken string
	Header      http.Header
}

type ReqOptionFunc func(option *ReqOption)

func WithHeader(header http.Header) ReqOptionFunc {
	return func(option *ReqOption) {
		option.Header = header
	}
}

func WithAccessToken(accessToken string) ReqOptionFunc {
	return func(option *ReqOption) {
		option.AccessToken = accessToken
	}
}
