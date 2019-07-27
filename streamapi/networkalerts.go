package streamapi

import (
	"context"
	. "net/http"
	"shodan/config"
	"shodan/net/http"
	"shodan/net/httputil"
	"strings"
)

type NetworkAlert interface {
	AllNetworkAlerts(ctx context.Context, outputType string) (string, error)
	FiltereByAlertID(ctx context.Context, outputType string, id string) (string, error)
}
type networkAlert struct {
	key    string
	configuration config.StreamNetworkAlerts
}

func (networkAlert *networkAlert) AllNetworkAlerts(ctx context.Context, outputType string) (string, error) {
	url := networkAlert.configuration.AllNetworkAlertsURL

	options := make(map[string]string)
	options[config.KEY] = networkAlert.key
	options[ "t"] = outputType

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (networkAlert *networkAlert) FiltereByAlertID(ctx context.Context, outputType string, id string) (string, error) {
	url := networkAlert.configuration.FiltereByAlertIDURL
	furl := strings.Replace(url, "{id}", id, -1)

	options := make(map[string]string)
	options[config.KEY] = networkAlert.key
	options[ "t"] = outputType

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}


