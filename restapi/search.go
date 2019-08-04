package restapi

import (
	"context"
	. "net/http"
	"strconv"
	"strings"
	"vighnesh.org/shodan/config"
	"vighnesh.org/shodan/net/http"
	"vighnesh.org/shodan/net/httputil"
)

const (
	SearchOptionHistory = "history"
	SearchOptionMinify  = "minify"
	SearchOptionQuery   = "query"
	SearchOptionFacets  = "facets"
	SearchOptionPage    = "page"
)

type Search interface {
	HostInfo(ctx context.Context, ip string, history, minify bool) (string, error)
	HostCount(ctx context.Context, query, facets string) (string, error)
	HostSearch(ctx context.Context, query, facets string, page int, minify bool) (string, error)
	HostSearchTokens(ctx context.Context, query string) (string, error)
	Ports(ctx context.Context) (string, error)
}

type search struct {
	key    string
	config config.Search
}

func (search *search) HostInfo(ctx context.Context, ip string, history, minify bool) (string, error) {
	url := search.config.HostInformationURL
	furl := strings.Replace(url, "{ip}", ip, -1)

	options := make(map[string]string)
	options[config.KEY] = search.key
	options[SearchOptionHistory] = strconv.FormatBool(history)
	options[SearchOptionMinify] = strconv.FormatBool(minify)

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)

}

func (search *search) HostCount(ctx context.Context, query, facets string) (string, error) {
	url := search.config.HostCountURL

	options := make(map[string]string)
	options[config.KEY] = search.key
	options[SearchOptionQuery] = query
	options[SearchOptionFacets] = facets

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (search *search) HostSearch(ctx context.Context, query, facets string, page int, minify bool) (string, error) {

	url := search.config.HostSearchURL

	options := make(map[string]string)
	options[config.KEY] = search.key
	options[SearchOptionQuery] = query
	options[SearchOptionFacets] = facets
	options[SearchOptionMinify] = strconv.FormatBool(minify)
	options[SearchOptionPage] = strconv.Itoa(page)

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)

}

func (search *search) HostSearchTokens(ctx context.Context, query string) (string, error) {
	url := search.config.HostSearchTokensURL

	options := make(map[string]string)
	options[config.KEY] = search.key
	options[SearchOptionQuery] = query

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (search *search) Ports(ctx context.Context) (string, error) {
	url := search.config.PortsURL

	options := make(map[string]string)
	options[config.KEY] = search.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}
