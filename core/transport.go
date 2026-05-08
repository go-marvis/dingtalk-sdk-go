package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
)

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
	for _, optf := range options {
		optf(option)
	}

	var (
		body []byte
		err  error
	)
	if apiReq.Body == nil {
		body = []byte{}
	} else {
		body, err = config.Serializer.Marshal([]any{apiReq.Body})
		if err != nil {
			return nil, err
		}
	}

	var values = make(url.Values, 0)
	for key := range apiReq.QueryParams {
		values.Set(key, apiReq.QueryParams.Get(key))
	}

	if apiReq.AccessTokenType != AccessTokenTypeNone {
		accessToken := option.AccessToken
		if accessToken == "" {
			accessToken, err = tokenManager.getAccessToken(ctx, config)
			if err != nil {
				return nil, err
			}
			values.Set("access_token", accessToken)
		}
	}

	path := config.BaseUrl + apiReq.ApiPath
	if query := values.Encode(); query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}

	if config.Debug {
		payload := body
		buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
		if err := json.Indent(buf, body, "", "  "); err == nil {
			payload = buf.Bytes()
		}
		log.Printf("[DEBUG] [API] %s %s payload:\n%s\n", apiReq.HttpMethod, path, payload)
	}

	req, err := http.NewRequestWithContext(ctx, apiReq.HttpMethod, path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set(contentTypeHeader, defaultContentType)

	var apiResp *ApiResp
	for range config.Retry {
		if config.Limiter != nil {
			if err = config.Limiter.Wait(ctx); err != nil {
				return nil, err
			}
		}

		apiResp, err = doSend(config.HttpClient, req)
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

			// TODO retry
		}

		err = nil
	}
	return apiResp, nil
}
