package core

import "time"

const (
	defaultContentType = "application/json; charset=utf-8"
)

const (
	contentTypeHeader = "Content-Type"
)

const (
	AppAccessTokenUrlPath string = "/gettoken"
)

const (
	appAccessTokenKeyPrefix = "app_access_token"
)

const expiryDelta = 3 * time.Minute

type AccessTokenType string

const (
	AccessTokenTypeNone AccessTokenType = "none_access_token"
)
