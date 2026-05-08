package core

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

var tokenManager = TokenManager{cache: cache}

type TokenManager struct {
	cache Cache
}

func (m *TokenManager) getAccessToken(ctx context.Context, config *Config) (string, error) {

	token, err := m.cache.Get(ctx, appAccessTokenKey(config.AppId))
	if err != nil {
		return "", err
	}

	if token == "" {
		token, err = m.getAppAccessTokenThenCache(ctx, config)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

type GetTokenResp struct {
	*ApiResp `json:"-"`
	CodeError
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (m *TokenManager) getAppAccessTokenThenCache(ctx context.Context, config *Config) (string, error) {
	req := NewApiReq()
	req.AccessTokenType = AccessTokenTypeNone
	req.ApiPath = AppAccessTokenUrlPath
	req.QueryParams.Add("appkey", config.AppKey)
	req.QueryParams.Add("appsecret", config.AppSecret)

	rawResp, err := Request(ctx, req, config)
	if err != nil {
		return "", err
	}

	resp := &GetTokenResp{}
	err = rawResp.UnmarshalBody(&resp, config)
	if err != nil {
		return "", err
	}
	expire := time.Duration(resp.ExpiresIn)*time.Second - expiryDelta
	err = m.cache.Set(ctx, appAccessTokenKey(config.AppId), resp.AccessToken, expire)
	if err != nil {
		slog.Warn("app appAccessToken save cache", "err", err)
	}

	return resp.AccessToken, nil
}
func appAccessTokenKey(appId string) string {
	return fmt.Sprintf("%s-%s", appAccessTokenKeyPrefix, appId)
}
