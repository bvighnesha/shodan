package restapi

import (
	"context"
	. "net/http"
	"shodan/config"
	"shodan/net/http"
	"shodan/net/httputil"
	"strconv"
)

const (
	DirectoryOptionPage  = "page"
	DirectoryOptionSort  = "sort"
	DirectoryOptionOrder = "order"
	DirectoryOptionQuery = "query"
	DirectoryOptionSize  = "size"
)

type Directory interface {
	Query(ctx context.Context, page int, sort, order string) (string, error)
	QuerySearch(ctx context.Context, query string, page int) (string, error)
	QueryTags(ctx context.Context, size int) (string, error)
}

type directory struct {
	key    string
	config config.Directory
}

func (directory *directory) Query(ctx context.Context, page int, sort, order string) (string, error) {
	url := directory.config.QueryURL

	options := make(map[string]string)

	options[config.KEY] = directory.key
	options[DirectoryOptionPage] = strconv.Itoa(page)
	options[DirectoryOptionSort] = sort
	options[DirectoryOptionOrder] = order

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (directory *directory) QuerySearch(ctx context.Context, query string, page int) (string, error) {
	url := directory.config.QuerySearchURL

	options := make(map[string]string)

	options[config.KEY] = directory.key
	options[DirectoryOptionQuery] = query
	options[DirectoryOptionPage] = strconv.Itoa(page)

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (directory *directory) QueryTags(ctx context.Context, size int) (string, error) {
	url := directory.config.QueryTagsURL

	options := make(map[string]string)

	options[config.KEY] = directory.key
	options[DirectoryOptionSize] = strconv.Itoa(size)

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)

}
