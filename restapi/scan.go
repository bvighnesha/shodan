package restapi

import (
	"context"
	. "net/http"
	"strings"
	"vighnesh.org/shodan/config"
	"vighnesh.org/shodan/net/http"
	"vighnesh.org/shodan/net/httputil"

)

type Scan interface {
	Protocols(ctx context.Context) (string, error)
	Scan(ctx context.Context, ips []string) (string, error)
	ScanInternet(ctx context.Context, port int, protocol string) (string, error)
	ScanStatus(ctx context.Context, id string) (string, error)
}

type scan struct {
	key    string
	config config.Scan
}

func (scan *scan) Protocols(ctx context.Context) (string, error) {
	url := scan.config.ProtocolsURL

	options := make(map[string]string)
	options[config.KEY] = scan.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (scan *scan) Scan(ctx context.Context, ips []string) (string, error) {
	panic("implement me")
}

func (scan *scan) ScanInternet(ctx context.Context, port int, protocol string) (string, error){
	panic("implement me")
}

func (scan *scan) ScanStatus(ctx context.Context, id string) (string, error) {
	url := scan.config.ScanStatusURL
	furl := strings.Replace(url, "{id}", id, -1)

	options := make(map[string]string)
	options[config.KEY] = scan.key

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}
