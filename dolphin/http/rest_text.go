package http

import (
	"net/http"
	"net/url"
)

type RestTextHandler struct {
	RestHandler
	Header  http.Header
	Cookies []*http.Cookie
}

func (handler RestTextHandler) ObtainUri() string {
	return handler.RestHandler.uri
}

func (handler RestTextHandler) ObtainParam() interface{} {
	return nil
}

func (handler RestTextHandler) ObtainValue() url.Values {
	return handler.values
}

func (handler RestTextHandler) ObtainRemoteServer() string {
	return handler.RestHandler.RemoteServer
}

func (handler RestTextHandler) ObtainHeader() http.Header {
	handler.Header.Add("Content-Type", "text/html")
	return handler.Header
}

func (handler RestTextHandler) ObtainCookies() []*http.Cookie {
	return handler.Cookies
}

func (handler *RestTextHandler) Post() (body []byte, err error) {
	return requestText(http.MethodPost, handler)
}

func (handler *RestTextHandler) Put() (body []byte, err error) {
	return requestText(http.MethodPut, handler)
}

func (handler *RestTextHandler) Delete() (body []byte, err error) {
	return requestText(http.MethodDelete, handler)
}

func (handler *RestTextHandler) Patch() (body []byte, err error) {
	return requestText(http.MethodPatch, handler)
}

func (handler *RestTextHandler) Options() (body []byte, err error) {
	return requestText(http.MethodOptions, handler)
}

func (handler *RestTextHandler) Head() (body []byte, err error) {
	return requestText(http.MethodHead, handler)
}

func (handler *RestTextHandler) Connect() (body []byte, err error) {
	return requestText(http.MethodConnect, handler)
}

func (handler *RestTextHandler) Trace() (body []byte, err error) {
	return requestText(http.MethodTrace, handler)
}

// Get 发送get请求
func (handler *RestTextHandler) Get() (body []byte, err error) {
	return get(handler)
}
