package core

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type AccessTokenResp struct {
	*ApiResp `json:"-"`
	CodeError
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expireIn"`
}

var tokenManager = TokenManager{cache: cache}

type TokenManager struct {
	cache Cache
}

func (m *TokenManager) getAppAccessToken(ctx context.Context, config *Config) (string, error) {
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

func (m *TokenManager) getCorpAccessToken(ctx context.Context, config *Config) (string, error) {
	token, err := m.cache.Get(ctx, corpAccessTokenKey(config.CorpId))
	if err != nil {
		return "", err
	}

	if token == "" {
		token, err = m.getCorpAccessTokenThenCache(ctx, config)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

type AccessTokenReq struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

// getAppAccessTokenThenCache 企业内部应用
func (m *TokenManager) getAppAccessTokenThenCache(ctx context.Context, config *Config) (string, error) {
	rawResp, err := Request(ctx, &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    AppAccessTokenUrlPath,
		Body: &AccessTokenReq{
			AppKey:    config.AppKey,
			AppSecret: config.AppSecret,
		},
		SupportedAccessTokenTypes: []AccessTokenType{AccessTokenTypeNone},
	}, config)
	if err != nil {
		return "", err
	}

	resp := &AccessTokenResp{}
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

type CorpAccessTokenReq struct {
	SuiteKey    string `json:"suiteKey"`
	SuiteSecret string `json:"suiteSecret"`
	AuthCorpId  string `json:"authCorpId"`
	SuiteTicket string `json:"suiteTicket"`
}

func (m *TokenManager) getCorpAccessTokenThenCache(ctx context.Context, config *Config) (string, error) {
	rawResp, err := Request(ctx, &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    CorpAccessTokenUrlPath,
		Body: &CorpAccessTokenReq{
			SuiteKey:    config.SuiteKey,
			SuiteSecret: config.SuiteSecret,
			AuthCorpId:  config.CorpId,
			SuiteTicket: config.SuiteTicket,
		},
		SupportedAccessTokenTypes: []AccessTokenType{AccessTokenTypeNone},
	}, config)

	resp := &AccessTokenResp{}
	err = rawResp.UnmarshalBody(&resp, config)
	if err != nil {
		return "", err
	}
	expire := time.Duration(resp.ExpiresIn)*time.Second - expiryDelta
	err = m.cache.Set(ctx, corpAccessTokenKey(config.AppId), resp.AccessToken, expire)
	if err != nil {
		slog.Warn("app corpAccessToken save cache", "err", err)
	}

	return resp.AccessToken, nil
}
