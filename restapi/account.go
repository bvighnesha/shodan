package restapi

import (
	"context"
	. "net/http"
	"vighnesha.in/shodan/config"
	"vighnesha.in/shodan/net/http"
	"vighnesha.in/shodan/net/httputil"
)

type Account interface {
	Profile(ctx context.Context) (string, error)
}

type account struct {
	key    string
	config config.Account
}

func (account *account) Profile(ctx context.Context) (string, error) {
	url := account.config.AccountProfileURL

	options := make(map[string]string)
	options[config.KEY] = account.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}
