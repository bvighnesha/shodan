package restapi

import (
	"context"
	. "net/http"
	"vighnesha.in/shodan/config"
	"vighnesha.in/shodan/net/http"
	"vighnesha.in/shodan/net/httputil"
)

type APIStatus interface {
	ApiInfo(ctx context.Context) (string, error)
}

type apiStatus struct {
	key    string
	config config.Apistatus
}

func (apiStatus *apiStatus) ApiInfo(ctx context.Context) (string, error) {
	url := apiStatus.config.APIPlanInformationURL

	options := make(map[string]string)
	options[config.KEY] = apiStatus.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}
