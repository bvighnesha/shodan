package restapi

import (
	"context"
	. "net/http"
	"shodan/config"
	"shodan/net/http"
	"shodan/net/httputil"
	"strconv"
	"strings"
)

const (
	User   = "user"
	Notify = "notify"
)

type Organization interface {
	Info(ctx context.Context) (string, error)
	AddNewMember(ctx context.Context, user string, notify bool) (string, error)
	RemoveMember(ctx context.Context, user string) (string, error)
}

type organization struct {
	key    string
	config config.Organization
}

func (org *organization) Info(ctx context.Context) (string, error) {
	url := org.config.InfoURL

	options := make(map[string]string)

	options[config.KEY] = org.key

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (org *organization) AddNewMember(ctx context.Context, user string, notify bool) (string, error) {
	url := org.config.AddNewMemberURL
	furl := strings.Replace(url, "{user}", user, -1)

	options := make(map[string]string)
	options[config.KEY] = org.key
	options[User] = user
	options[Notify] = strconv.FormatBool(notify)

	response, e := http.Do(ctx, MethodPut, furl, options)
	return httputil.Response(response, e)
}

func (org *organization) RemoveMember(ctx context.Context, user string) (string, error) {
	url := org.config.RemoveMemberURL
	furl := strings.Replace(url, "{user}", user, -1)

	options := make(map[string]string)
	options[config.KEY] = org.key
	options[User] = user

	response, e := http.Do(ctx, MethodPut, furl, options)
	return httputil.Response(response, e)
}
