package net

import (
	"io/ioutil"
	"net/http"
)

type HttpRequester interface {
	Get(uri string, headers map[string]string) (*HttpResponse, error)
}

type DefaultHttpRequester struct {
	Client *http.Client
}

func (d *DefaultHttpRequester) Get(uri string, headers map[string]string) (*HttpResponse, error) {
	client := d.Client
	request, e := http.NewRequest("GET", uri, nil)
	if e != nil {
		return nil, e
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	response, e := client.Do(request)
	if e != nil {
		return nil, e
	}
	defer response.Body.Close()
	httpResponse := &HttpResponse{}
	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, e
	}
	httpResponse.Body = bytes
	httpResponse.Headers = response.Header
	httpResponse.Uri = uri
	httpResponse.Code = response.StatusCode
	return httpResponse, nil
}

func (d *DefaultHttpRequester) Post(uri string, bodyType string, body interface{}, headers map[string][]string) (http.Response, error) {
	panic("implement me")
}

type HttpResponse struct {
	Uri     string
	Code    int
	Headers http.Header
	Body    []byte
}
