package core

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Config struct {
	OapiUrl string
	BaseUrl string // 服务地址

	AppId       string // 第三方个人应用
	AppKey      string // 企业内部应用
	AppSecret   string
	CorpId      string // 第三方应用授权企业
	SuiteKey    string // 第三方企业应用
	SuiteSecret string
	SuiteTicket string

	Debug bool

	Retry   int           // 重试次数
	Limiter *rate.Limiter // 请求限速

	HttpClient HttpClient
	ReqTimeout time.Duration

	Serializer Serializer
}

func Init(config *Config) {
	if config.HttpClient == nil {
		if config.ReqTimeout == 0 {
			config.HttpClient = http.DefaultClient
		} else {
			config.HttpClient = &http.Client{Timeout: config.ReqTimeout}
		}
	}

	if config.Serializer == nil {
		config.Serializer = &defaultSerializer{}
	}
}
