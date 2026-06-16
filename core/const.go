package core

import (
	"fmt"
	"time"
)

const (
	defaultContentType = "application/json;charset=utf-8"
)

const (
	OapiUrl = "https://oapi.dingtalk.com"
	BaseUrl = "https://api.dingtalk.com"
)

const (
	// 企业内部应用
	AppAccessTokenUrlPath = "/v1.0/oauth2/accessToken"
	// 第三方应用授权企业
	CorpAccessTokenUrlPath = "/oauth2/corpAccessToken"
	// 第三方个人应用
	UserAccessTokenUrlPath = "/sns/gettoken"
	// 第三方企业应用
	SuiteAccessTokenUrlPath = "/v1.0/oauth2/suiteAccessToken"
)

func appAccessTokenKey(appId string) string {
	return fmt.Sprintf("%s-%s", "app_access_token", appId)
}

func corpAccessTokenKey(corpId string) string {
	return fmt.Sprintf("%s-%s", "corp_access_token", corpId)
}

const expiryDelta = 3 * time.Minute

type AccessTokenType string

const (
	AccessTokenTypeNone  AccessTokenType = "none_access_token"
	AccessTokenTypeApp   AccessTokenType = "app_access_token"
	AccessTokenTypeCorp  AccessTokenType = "corp_access_token"
	AccessTokenTypeUser  AccessTokenType = "user_access_token"
	AccessTokenTypeSuite AccessTokenType = "suite_access_token"
)
