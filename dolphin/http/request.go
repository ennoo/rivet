package http

import (
	"bytes"
	"encoding/json"
	"github.com/ennoo/rivet/common/log"
	"net/http"
	"net/url"
	"path/filepath"
)

type Request interface {
	Post() (body []byte, err error)

	Put() (body []byte, err error)

	Delete() (body []byte, err error)

	Patch() (body []byte, err error)

	Options() (body []byte, err error)

	Head() (body []byte, err error)

	Connect() (body []byte, err error)

	Trace() (body []byte, err error)

	Get() (body []byte, err error)
}

type Handler interface {
	ObtainUri() string

	ObtainParam() interface{}

	ObtainValue() url.Values

	ObtainRemoteServer() string

	ObtainHeader() http.Header

	ObtainCookies() []*http.Cookie
}

func getFullUri(handler Handler) string {
	url := filepath.Join(handler.ObtainRemoteServer(), "/", handler.ObtainUri())
	url = filepath.ToSlash(url)
	return url
}

func requestJson(method string, handler Handler) ([]byte, error) {
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(handler.ObtainParam())
	log.Debug("body:", string(jsonByte))
	req, err := http.NewRequest(method, getFullUri(handler), bytes.NewReader(jsonByte))
	if nil != err {
		return nil, err
	}
	req.Cookies() = handler.ObtainCookies()
	req.Header = handler.ObtainHeader()
	return response(req)
}

func requestText(method string, handler Handler) ([]byte, error) {
	req, err := http.NewRequest(method, getFullUri(handler), bytes.NewBufferString(handler.ObtainValue().Encode()))
	if nil != err {
		return nil, err
	}
	req.Cookies() = handler.ObtainCookies()
	req.Header = handler.ObtainHeader()
	return response(req)
}

// Get 发送get请求
func get(handler Handler) (body []byte, err error) {
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(handler.ObtainParam())
	log.Debug("body:", string(jsonByte))
	req, err := http.NewRequest(http.MethodGet, getFullUri(handler), nil)
	if nil != err {
		return nil, err
	}
	req.Cookies() = handler.ObtainCookies()
	req.Header = handler.ObtainHeader()
	return response(req)
}
