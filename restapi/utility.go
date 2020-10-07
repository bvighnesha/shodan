package restapi

import (
	"context"
	. "net/http"
	"vighnesha.in/shodan/config"
	"vighnesha.in/shodan/net/http"
	"vighnesha.in/shodan/net/httputil"
)

type Utility interface {
	HTTPHeaders(ctx context.Context) (string, error)
	MyIP(ctx context.Context) (string, error)
}

type utility struct {
	key    string
	config config.Utility
}

func (utility *utility) HTTPHeaders(ctx context.Context) (string, error) {
	url := utility.config.HTTPHeadersURL

	options := make(map[string]string)
	options[config.KEY] = utility.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (utility *utility) MyIP(ctx context.Context) (string, error) {
	url := utility.config.MyIPAddressURL

	options := make(map[string]string)
	options[config.KEY] = utility.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}
