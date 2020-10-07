package restapi

import (
	"context"
	. "net/http"
	"vighnesha.in/shodan/config"
	"vighnesha.in/shodan/net/http"
	"vighnesha.in/shodan/net/httputil"
)

const DataOptionsDataset = "dataset"

type Data interface {
	Data(ctx context.Context) (string, error)
	DataSet(ctx context.Context, dataset string) (string, error)
}

type data struct {
	key    string
	config config.Data
}

func (data *data) Data(ctx context.Context) (string, error) {
	url := data.config.DataURL

	options := make(map[string]string)

	options[config.KEY] = data.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (data *data) DataSet(ctx context.Context, dataset string) (string, error) {
	url := data.config.DatasetURL

	options := make(map[string]string)

	options[config.KEY] = data.key
	options[DataOptionsDataset] = dataset

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}
