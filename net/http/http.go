package http

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"vighnesha.in/shodan/net"
)

const UserAgentKey = "User-Agent"
const ContentTypeKey = "Content-Type"

func Do(ctx context.Context, method, url string, options map[string]string) (*http.Response, error) {
	request, e := http.NewRequest(method, url, nil)
	if e != nil {
		return nil, e
	}
	request.Header.Add(UserAgentKey, net.UserAgent)
	query := request.URL.Query()

	for key, value := range options {
		if key != "" && value != "" {
			query.Add(key, value)
		}

	}

	request.URL.RawQuery = Encode(query)

	fmt.Println(Encode(query))
	return Execute(ctx, request)
}

func Execute(ctx context.Context, request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	if ctx != nil {
		request = request.WithContext(ctx)
	}

	return client.Do(request)

}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func Encode(values map[string][]string) string {
	if values == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := values[k]

		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}
