package restapi

import (
	"context"
	. "net/http"
	"vighnesh.org/shodan/config"
	"vighnesh.org/shodan/net/http"
	"vighnesh.org/shodan/net/httputil"

	"strings"
)

type Network interface {
	Alert(ctx context.Context, name string, filters interface{}, filtersip []string, expires int)
	AlertInfoById(ctx context.Context, id string) (string, error)
	Delete(ctx context.Context, id string) (string, error)
	AlertInfo(ctx context.Context) (string, error)
	Triggers(ctx context.Context) (string, error)
	EnableTrigger(ctx context.Context, id, trigger string) (string, error)
	DisableTrigger(ctx context.Context, id, trigger string) (string, error)
	AddToWhitelist(ctx context.Context, id, trigger, service string) (string, error)
	RemoveFromWhitelist(ctx context.Context, id, trigger, service string) (string, error)
}

type network struct {
	key    string
	config config.NetworkAlerts
}

func (network *network) Alert(ctx context.Context, name string, filters interface{}, filtersip []string, expires int) {
	panic("implement me")
}

func (network *network) AlertInfoById(ctx context.Context, id string) (string, error) {
	url := network.config.AlertIDInfoURL
	furl := strings.Replace(url, "{id}", id, -1)

	options := make(map[string]string)
	options[config.KEY] = network.key

	response, e := http.Do(ctx, MethodGet, furl, options)
	return httputil.Response(response, e)
}

func (network *network) Delete(ctx context.Context, id string) (string, error) {
	url := network.config.DeleteAlertURL
	furl := strings.Replace(url, "{id}", id, -1)

	options := make(map[string]string)
	options[config.KEY] = network.key

	response, e := http.Do(ctx, MethodDelete, furl, options)
	return httputil.Response(response, e)
}

func (network *network) AlertInfo(ctx context.Context) (string, error) {
	url := network.config.AlertInfoURL

	options := make(map[string]string)
	options[config.KEY] = network.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (network *network) Triggers(ctx context.Context) (string, error) {
	url := network.config.AlertTriggersURL

	options := make(map[string]string)
	options[config.KEY] = network.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (network *network) EnableTrigger(ctx context.Context, id, trigger string) (string, error) {
	url := network.config.AlertTriggersURL
	idReplaced := strings.Replace(url, "{id}", id, -1)
	furl := strings.Replace(idReplaced, "{trigger}", trigger, -1)

	options := make(map[string]string)
	options[config.KEY] = network.key

	response, e := http.Do(ctx, MethodPut, furl, options)
	return httputil.Response(response, e)
}

func (network *network) DisableTrigger(ctx context.Context, id, trigger string) (string, error) {
	url := network.config.DisableTriggerURL
	idReplaced := strings.Replace(url, "{id}", id, -1)
	furl := strings.Replace(idReplaced, "{trigger}", trigger, -1)

	options := make(map[string]string)
	options[config.KEY] = network.key

	response, e := http.Do(ctx, MethodDelete, furl, options)
	return httputil.Response(response, e)
}

func (network *network) AddToWhitelist(ctx context.Context, id, trigger, service string) (string, error) {
	url := network.config.AddTriggerToWhitelistURL
	idReplaced := strings.Replace(url, "{id}", id, -1)
	triggerReplaced := strings.Replace(idReplaced, "{trigger}", trigger, -1)
	furl := strings.Replace(triggerReplaced, "{service}", trigger, -1)

	options := make(map[string]string)
	options[config.KEY] = network.key

	response, e := http.Do(ctx, MethodPut, furl, options)
	return httputil.Response(response, e)
}

func (network *network) RemoveFromWhitelist(ctx context.Context, id, trigger, service string) (string, error) {
	url := network.config.RemoveFromWhitelistURL
	idReplaced := strings.Replace(url, "{id}", id, -1)
	triggerReplaced := strings.Replace(idReplaced, "{trigger}", trigger, -1)
	furl := strings.Replace(triggerReplaced, "{service}", trigger, -1)

	options := make(map[string]string)
	options[config.KEY] = network.key

	response, e := http.Do(ctx, MethodDelete, furl, options)
	return httputil.Response(response, e)
}
