package streamapi

import (
	"context"
	. "net/http"
	"vighnesh.org/shodan/config"
	"vighnesh.org/shodan/net/http"
	"vighnesh.org/shodan/net/httputil"

	"strings"
)

type Data interface {
	Banners(ctx context.Context, outputType string) (string, error)
	FiltereByASN(ctx context.Context, outputType string, asn string) (string, error)
	FiltereByountry(ctx context.Context, outputType string, countries string) (string, error)
	FiltereByPorts(ctx context.Context, outputType string, ports string) (string, error)
}

type data struct {
	key    string
	configuration config.StreamData
}

func (data *data) Banners(ctx context.Context, outputType string) (string, error) {
	url := data.configuration.BannersURL

	options := make(map[string]string)
	options[config.KEY] = data.key
	options[ "t"] = outputType

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (data *data) FiltereByASN(ctx context.Context, outputType string, asn string) (string, error) {
	url := data.configuration.FiltereByASNURL
	furl := strings.Replace(url, "{asn}", asn, -1)

	options := make(map[string]string)
	options[config.KEY] = data.key
	options[ "t"] = outputType

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}

func (data *data) FiltereByountry(ctx context.Context, outputType string, countries string) (string, error) {
	url := data.configuration.FiltereByCountryURL
	furl := strings.Replace(url, "{countries}", countries, -1)

	options := make(map[string]string)
	options[config.KEY] = data.key
	options[ "t"] = outputType

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}

func (data *data) FiltereByPorts(ctx context.Context, outputType string, ports string) (string, error) {
	url := data.configuration.FiltereByPortsURL
	furl := strings.Replace(url, "{ports}", ports, -1)

	options := make(map[string]string)
	options[config.KEY] = data.key
	options[ "t"] = outputType

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}

