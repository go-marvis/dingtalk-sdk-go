package core

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

var translator ReqTranslator

func determineTokenType(accessTokenTypes []AccessTokenType, option *RequestOption, config *Config) AccessTokenType {
	if option.AppAccessToken != "" {
		return AccessTokenTypeApp
	}
	if option.CorpAccessToken != "" {
		return AccessTokenTypeCorp
	}
	if option.SuiteAccessToken != "" {
		return AccessTokenTypeSuite
	}

	accessTokenType := accessTokenTypes[0]
	tokenSet := make(map[AccessTokenType]struct{})
	for _, t := range accessTokenTypes {
		tokenSet[t] = struct{}{}
	}

	if config.AppKey != "" {
		if _, ok := tokenSet[AccessTokenTypeApp]; ok {
			return AccessTokenTypeApp
		}
	}
	if config.CorpId != "" {
		if _, ok := tokenSet[AccessTokenTypeCorp]; ok {
			return AccessTokenTypeCorp
		}
	}
	if config.AppId != "" {
		if _, ok := tokenSet[AccessTokenTypeUser]; ok {
			return AccessTokenTypeUser
		}
	}
	if config.SuiteKey != "" {
		if _, ok := tokenSet[AccessTokenTypeSuite]; ok {
			return AccessTokenTypeSuite
		}
	}

	return accessTokenType
}

func doSend(client HttpClient, req *http.Request) (*ApiResp, error) {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ApiResp{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		RawBody:    body,
	}, nil
}

func readResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Request(ctx context.Context, apiReq *ApiReq, config *Config, options ...RequestOptionFunc) (*ApiResp, error) {
	option := &RequestOption{}
	for _, fn := range options {
		fn(option)
	}

	accessTokenType := determineTokenType(apiReq.SupportedAccessTokenTypes, option, config)
	httpReq, err := translator.translate(ctx, apiReq, accessTokenType, config, option)
	if err != nil {
		return nil, err
	}

	var apiResp *ApiResp
	if config.Limiter != nil {
		if err = config.Limiter.Wait(ctx); err != nil {
			return nil, err
		}
	}

	apiResp, err = doSend(config.HttpClient, httpReq)
	if err != nil {
		return apiResp, err
	}

	if apiResp.StatusCode != http.StatusOK { // http err
		slog.Error("http error", "status", apiResp.StatusCode)
		return apiResp, fmt.Errorf("http error, status=%d", apiResp.StatusCode)
	}

	codeError := &CodeError{}
	if err = config.Serializer.Unmarshal(apiResp.RawBody, codeError); err != nil {
		slog.Error("unmarshal error", "err", err)
		return apiResp, err
	}

	if codeError.ErrCode != 0 { // api err
		slog.Error("api error", "errcode", codeError.ErrCode, "errmsg", codeError.ErrMsg)
		err = fmt.Errorf("api error, code=%v, err=%v", codeError.ErrCode, codeError.ErrMsg)
		return apiResp, err
	}

	return apiResp, err
}
