package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"shodan/net"
)

const UserAgentKey = "User-Agent"
const ContentTypeKey = "Content-Type"

func Do(ctx context.Context, method, url string, options map[string]string) (*http.Response, error) {
	request, e := http.NewRequest(method, url, nil)
	if e != nil {
		return nil, e
	}
	request.Header.Add(UserAgentKey, net.USER_AGENT)
	query := request.URL.Query()

	for key, value := range options {
		if key != "" && value != "" {
			query.Add(key, value)
		}

	}

	request.URL.RawQuery = query.Encode()

	fmt.Println(query.Encode())
	return Execute(ctx, request)
}

func DoPost(ctx context.Context, me *net.Entity, url string) (*http.Response, error) {
	request, e := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(me.Build()))
	if e != nil {
		return nil, e
	}

	request.Header.Set(ContentTypeKey, net.CONTENT_TYPE)
	request.Header.Add(UserAgentKey, net.USER_AGENT)
	return Execute(ctx, request)
}

func Execute(ctx context.Context, request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	if ctx != nil {
		request = request.WithContext(ctx)
	}

	return client.Do(request)

}
