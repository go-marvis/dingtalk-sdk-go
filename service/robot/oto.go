package robot

import (
	"context"
	"net/http"

	"github.com/go-marvis/dingtalk-sdk-go/core"
)

type oto struct {
	config *core.Config
}

func (s *oto) BatchSend(ctx context.Context, req *BatchSendOTOReq, options ...core.RequestOptionFunc) (*BatchSendOTOResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.ApiPath = "/v1.0/robot/oToMessages/batchSend"
	apiReq.SupportedAccessTokenTypes = []core.AccessTokenType{
		core.AccessTokenTypeApp,
		core.AccessTokenTypeCorp,
	}

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	resp := &BatchSendOTOResp{ApiResp: apiResp}
	if err != nil {
		return resp, err
	}

	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
