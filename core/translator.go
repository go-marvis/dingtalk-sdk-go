package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type ReqTranslator struct {
}

func authorizationToRequest(req *ApiReq, option *RequestOption, token string) {
	if strings.Index(req.ApiPath, OapiUrl) == 0 {
		req.QueryParams.Set("access_token", token)
	} else {
		option.Header.Set("x-acs-dingtalk-access-token", token)
	}
}

func (translator *ReqTranslator) translate(ctx context.Context, req *ApiReq, accessTokenType AccessTokenType, config *Config, option *RequestOption) (*http.Request, error) {
	body := req.Body

	if _, ok := body.(*Formdata); ok {
		option.FileUpload = true
	}

	contentType, rawBody, err := translator.payload(body, config.Serializer)
	if err != nil {
		return nil, err
	}

	switch accessTokenType {
	case AccessTokenTypeApp:
		accessToken := option.AppAccessToken
		if accessToken == "" {
			accessToken, err = tokenManager.getAppAccessToken(ctx, config)
			if err != nil {
				return nil, err
			}
		}
		authorizationToRequest(req, option, accessToken)

	case AccessTokenTypeCorp:
		accessToken := option.CorpAccessToken
		if accessToken == "" {
			accessToken, err = tokenManager.getCorpAccessToken(ctx, config)
			if err != nil {
				return nil, err
			}
		}
		authorizationToRequest(req, option, accessToken)
	}

	var pathSegs []string
	for _, p := range strings.Split(req.ApiPath, "/") {
		if strings.Index(p, ":") == 0 {
			varName := p[1:]
			v, ok := req.PathParams[varName]
			if !ok {
				return nil, fmt.Errorf("http path:%s, name: %s, not found value", req.ApiPath, varName)
			}
			val := fmt.Sprint(v)
			if val == "" {
				return nil, fmt.Errorf("http path:%s, name: %s, value is empty", req.ApiPath, varName)
			}
			val = url.PathEscape(val)
			pathSegs = append(pathSegs, val)
			continue
		}
		pathSegs = append(pathSegs, p)
	}

	newPath := strings.Join(pathSegs, "/")
	if strings.Index(newPath, "http") != 0 {
		newPath = fmt.Sprintf("%s%s", config.BaseUrl, newPath)
	}
	if query := req.QueryParams.Encode(); query != "" {
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

	// create request
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
		httpReq.Header.Set("Content-Type", contentType)
	}
	return httpReq, nil
}

func (translator *ReqTranslator) payload(body interface{}, serializer Serializer) (string, []byte, error) {
	if fd, ok := body.(*Formdata); ok {
		return fd.content()
	}
	contentType := defaultContentType
	if body == nil {
		return contentType, nil, nil
	}

	bs, err := serializer.Marshal(body)
	return contentType, bs, err
}

type Formdata struct {
	fields map[string]interface{}
	data   *struct {
		content     []byte
		contentType string
	}
}

func (fd *Formdata) AddField(field string, val interface{}) *Formdata {
	if fd.fields == nil {
		fd.fields = map[string]interface{}{}
	}
	fd.fields[field] = val
	return fd
}

func (fd *Formdata) AddFile(field string, r io.Reader) *Formdata {
	return fd.AddField(field, r)
}

func (fd *Formdata) content() (string, []byte, error) {
	if fd.data != nil {
		return fd.data.contentType, fd.data.content, nil
	}
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for key, val := range fd.fields {
		if r, ok := val.(io.Reader); ok {
			filename := "unknown-file"
			if f, ok := val.(*os.File); ok {
				filename = filepath.Base(f.Name())
			}
			part, err := writer.CreateFormFile(key, filename)
			if err != nil {
				return "", nil, err
			}
			_, err = io.Copy(part, r)
			if err != nil {
				return "", nil, err
			}
			continue
		}
		err := writer.WriteField(key, fmt.Sprint(val))
		if err != nil {
			return "", nil, err
		}
	}
	contentType := writer.FormDataContentType()
	err := writer.Close()
	if err != nil {
		return "", nil, err
	}
	fd.data = &struct {
		content     []byte
		contentType string
	}{content: buf.Bytes(), contentType: contentType}
	return fd.data.contentType, fd.data.content, nil
}
