package restapi

import (
	"context"
	. "net/http"
	"vighnesh.org/shodan/config"
	"vighnesh.org/shodan/net/http"
	"vighnesh.org/shodan/net/httputil"

	"strings"
)

const Hostnames = "hostnames"
const IPs = "ips"

type DNS interface {
	DomainInfo(ctx context.Context, domain string) (string, error)
	Lookup(ctx context.Context, hostnames string) (string, error)
	ReverseLookup(ctx context.Context, ips string) (string, error)
}

type dns struct {
	key    string
	config config.DNS
}

func (dns *dns) DomainInfo(ctx context.Context, domain string) (string, error) {
	url := dns.config.DomainInfoURL

	options := make(map[string]string)
	options[config.KEY] = dns.key

	url = strings.Replace(url, "{domain}", domain, -1)

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (dns *dns) Lookup(ctx context.Context, hostnames string) (string, error) {
	url := dns.config.DNSLookupURL

	options := make(map[string]string)

	options[config.KEY] = dns.key
	options[Hostnames] = hostnames

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}

func (dns *dns) ReverseLookup(ctx context.Context, ips string) (string, error) {
	url := dns.config.ReverseDNSLookupURL

	options := make(map[string]string)
	options[config.KEY] = dns.key
	options[IPs] = ips

	response, e := http.Do(ctx, MethodGet, url, options)
	return httputil.Response(response, e)
}
