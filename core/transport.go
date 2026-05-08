package core

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

var translator ReqTranslator

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

func Request(ctx context.Context, apiReq *ApiReq, config *Config, options ...ReqOptionFunc) (*ApiResp, error) {
	option := &ReqOption{}
	for _, fn := range options {
		fn(option)
	}

	httpReq, err := translator.translate(ctx, apiReq, apiReq.AccessTokenType, config, option)
	if err != nil {
		return nil, err
	}

	var apiResp *ApiResp
	for range config.Retry {
		if config.Limiter != nil {
			if err = config.Limiter.Wait(ctx); err != nil {
				return nil, err
			}
		}

		apiResp, err = doSend(config.HttpClient, httpReq)
		if apiResp.StatusCode != http.StatusOK {
			slog.Error("http error", "status", apiResp.StatusCode)
			return apiResp, fmt.Errorf("http error, status=%d", apiResp.StatusCode)
		}

		codeError := &CodeError{}
		if err = config.Serializer.Unmarshal(apiResp.RawBody, codeError); err != nil {
			slog.Error("unmarshal error", "err", err)
			return apiResp, err
		}

		if codeError.ErrCode != 0 {
			slog.Error("api error", "errcode", codeError.ErrCode, "errmsg", codeError.ErrMsg)
			err = fmt.Errorf("api error, code=%v, err=%v", codeError.ErrCode, codeError.ErrMsg)
			return apiResp, err
		}

		err = nil
	}

	return apiResp, err
}
