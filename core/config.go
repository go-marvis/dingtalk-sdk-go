package core

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Config struct {
	BaseUrl string // 服务地址
	AppId   string

	AppKey    string // 接口账号
	AppSecret string
	Debug     bool

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
