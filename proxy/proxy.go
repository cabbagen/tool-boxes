package proxy

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type handler interface {
	handle(w http.ResponseWriter, req *http.Request)
}

type cproxy struct {
	BeforeRequestHandles []handler
	AfterResponseHandles []handler
}

func NewCproxy() cproxy {
	return cproxy{
		BeforeRequestHandles: []handler{},
		AfterResponseHandles: []handler{},
	}
}

func (cp cproxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	proxyResponse, proxyResponseError := NewHttpProxyRequest(cp, w, req)

	defer proxyResponse.Body.Close()

	if proxyResponseError != nil {
		log.Fatal("代理过程发生错误", proxyResponseError)
	}

	content, readError := ioutil.ReadAll(proxyResponse.Body)

	if readError != nil {
		log.Fatal("数据接受出现错误: ", readError)
	}

	AddHeaders(w, proxyResponse.Header)

	io.WriteString(w, string(content[0:]))
}

func NewHttpProxyRequest(cp cproxy, w http.ResponseWriter, req *http.Request) (*http.Response, error) {
	proxyClient := &http.Client{}

	fmt.Println("请求地址：", req.URL.String())

	proxyRequest, proxyRequestError := http.NewRequest(req.Method, req.URL.String(), req.Body)

	if proxyRequestError != nil {
		return nil, errors.New("代理请求发生错误")
	}

	// handle beforeRequestHandles
	for _, beforeHandle := range cp.BeforeRequestHandles {
		beforeHandle.handle(w, req)
	}

	proxyResponse, proxyResponseError := proxyClient.Do(proxyRequest)

	if proxyResponseError != nil {
		return nil, errors.New("代理相应发生错误")
	}

	// handle afterResponseHandles
	for _, afterHandle := range cp.AfterResponseHandles {
		afterHandle.handle(w, req)
	}

	return proxyResponse, nil
}

func AddHeaders(responseWriter http.ResponseWriter, header http.Header) {
	contentType := header.Get("Content-Type")

	responseWriter.Header().Set("Content-Type", contentType)
	responseWriter.Header().Set("proxy", "http-proxy")
}

func (cp cproxy) registerBeforeRequestHandle(handles []handler) {
	cp.BeforeRequestHandles = append(cp.BeforeRequestHandles, handles...)
}

func (cp cproxy) registerAfterResponseHandle(handles []handler) {
	cp.AfterResponseHandles = append(cp.AfterResponseHandles, handles...)
}
