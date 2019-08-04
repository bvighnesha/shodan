package restapi

import (
	"context"
	. "net/http"
	"vighnesh.org/shodan/config"
	"vighnesh.org/shodan/net/http"
	"vighnesh.org/shodan/net/httputil"

	"strings"
)

type Experimental interface {
	Honeyscore(ctx context.Context, ip string) (string, error)
}

type experimental struct {
	key    string
	config config.Experimental
}

func (experimental *experimental) Honeyscore(ctx context.Context, ip string) (string, error) {
	url := experimental.config.HoneyscoreURL
	furl := strings.Replace(url, "{ip}", ip, -1)

	options := make(map[string]string)
	options[config.KEY] = experimental.key

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}
