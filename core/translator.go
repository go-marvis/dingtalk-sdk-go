package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ReqTranslator struct {
}

func (translator *ReqTranslator) translate(ctx context.Context, req *ApiReq, accessTokenType AccessTokenType, config *Config, option *ReqOption) (*http.Request, error) {

	body := req.Body
	contentType, rawBody, err := translator.payload(body, config.Serializer)
	if err != nil {
		return nil, err
	}

	params := req.QueryParams
	if req.AccessTokenType != AccessTokenTypeNone {
		accessToken := option.AccessToken
		if accessToken == "" {
			accessToken, err = tokenManager.getAccessToken(ctx, config)
			if err != nil {
				return nil, err
			}
			params.Set("access_token", accessToken)
		}
	}

	newPath := config.BaseUrl + req.ApiPath
	if query := params.Encode(); query != "" {
		newPath = fmt.Sprintf("%s?%s", newPath, query)
	}

	if config.Debug {
		payload := rawBody
		buf := bytes.NewBuffer(make([]byte, 0, len(rawBody)+1024))
		if err := json.Indent(buf, rawBody, "", "  "); err == nil {
			payload = buf.Bytes()
		}
		log.Printf("[DEBUG] [API] %s %s payload:\n%s\n", req.HttpMethod, newPath, payload)
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.HttpMethod, newPath, bytes.NewBuffer(rawBody))
	if err != nil {
		return nil, err
	}

	for key, vals := range option.Header {
		for _, val := range vals {
			httpReq.Header.Add(key, val)
		}
	}

	if contentType != "" {
		httpReq.Header.Set(contentTypeHeader, contentType)
	}

	return httpReq, nil
}

func (translator *ReqTranslator) payload(body interface{}, serializer Serializer) (string, []byte, error) {
	contentType := defaultContentType
	if body == nil {
		return contentType, nil, nil
	}

	bs, err := serializer.Marshal(body)
	return contentType, bs, err
}
