package robot

import "github.com/go-marvis/dingtalk-sdk-go/core"

const (
	MsgKeyText       = "sampleText"
	MsgKeyMarkdown   = "sampleMarkdown"
	MsgKeyImage      = "sampleImageMsg"
	MsgKeyActionCard = "sampleActionCard"
	MsgKeyFeedCard   = ""
	MsgKeyLink       = "sampleLink"
	MsgKeyAudio      = "sampleAudio"
	MsgKeyFile       = "sampleFile"
	MsgKeyVideo      = "sampleVideo"
)

type MsgText struct {
	Content string `json:"content"`
}

type MsgMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type MsgImage struct {
	PhotoURL string `json:"photoURL"`
}

type MsgLink struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type MsgActionCard struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	SingleTitle string `json:"singleTitle"`
	SingleURL   string `json:"singleURL"`
}

type MsgAudio struct {
	MediaId  string `json:"mediaId"`
	Duration string `json:"duration"`
}

type MsgFile struct {
	MediaId  string `json:"mediaId"`
	FileId   string `json:"fileId"`
	FileType string `json:"fileType"`
}

type MsgVideo struct {
	VideoMediaId string `json:"videoMediaId"`
	VideoType    string `json:"videoType"`
	PicMediaId   string `json:"picMediaId"`
	Duration     string `json:"duration"`
}

type request struct {
	apiReq *core.ApiReq
}

type requestBuilder struct {
	apiReq *core.ApiReq
}
