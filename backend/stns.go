package backend

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/STNS/libstns-go/libstns"
)

// https://github.com/STNS/libnss/blob/4e91b751d643b58b32debaef287c95ce2b81c73c/stns.h#L25
var stnsClientConfFile = "/etc/stns/client/stns.conf"

type STNS struct {
	client *libstns.STNS
}

// NewSTNS ...
func NewSTNS() (*STNS, error) {
	endpoint := os.Getenv("STNS_API_ENDPOINT")
	options := &libstns.Options{}
	if _, err := os.Stat(stnsClientConfFile); err == nil {
		b, err := ioutil.ReadFile(stnsClientConfFile)
		if err != nil {
			return nil, err
		}
		c := &STNSClientConfig{}
		if _, err := toml.Decode(string(b), &c); err != nil {
			return nil, err
		}
		if endpoint == "" {
			endpoint = c.ApiEndpoint
		}
		options.AuthToken = c.AuthToken
		options.User = c.User
		options.Password = c.Password
		options.SkipSSLVerify = !c.SSLVerify
		options.HttpProxy = c.HTTPProxy
		options.RequestTimeout = c.RequestTimeout
		options.RequestRetry = c.RequestRetry
		if c.TLS != nil {
			options.TLS.CA = c.TLS.CA
			options.TLS.Cert = c.TLS.Cert
			options.TLS.Key = c.TLS.Key
		}
	}
	client, err := libstns.NewSTNS(endpoint, options)
	if err != nil {
		return nil, err
	}
	return &STNS{
		client: client,
	}, nil
}

// Keys ...
func (b *STNS) Keys(ctx context.Context, options ...Option) ([]string, error) {
	keys := []string{}
	c := &Config{}
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	for _, u := range c.Users {
		user, err := b.client.GetUserByName(u)
		if err != nil {
			return nil, err
		}
		keys = append(keys, user.Keys...)
	}

	for _, g := range c.Groups {
		group, err := b.client.GetGroupByName(g)
		if err != nil {
			return nil, err
		}
		for _, u := range group.Users {
			user, err := b.client.GetUserByName(u)
			if err != nil {
				return nil, err
			}
			keys = append(keys, user.Keys...)
		}
	}

	return unique(keys), nil
}
