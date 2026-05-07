package core

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ApiResp struct {
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	RawBody    []byte      `json:"-"`
}

func (resp ApiResp) UnmarshalBody(val any, config *Config) error {
	if config.Debug {
		payload := resp.RawBody
		buf := bytes.NewBuffer(make([]byte, 0, len(resp.RawBody)+1024))
		if err := json.Indent(buf, resp.RawBody, "", "  "); err == nil {
			payload = buf.Bytes()
		}
		log.Printf("[DEBUG] [API] response:\n%s\n", payload)
	}

	return config.Serializer.Unmarshal(resp.RawBody, val)
}

type CodeError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
