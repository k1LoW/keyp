package backend

import (
	"context"
	"fmt"
)

var Backends = []string{
	"github",
	"stns",
}

type Config struct {
	Users  []string
	Groups []string
}

type Option func(*Config) error

// Users ...
func Users(users []string) (Option, error) {
	return func(c *Config) error {
		c.Users = append(c.Users, users...)
		return nil
	}, nil
}

// Groups ...
func Groups(groups []string) (Option, error) {
	return func(c *Config) error {
		c.Groups = append(c.Groups, groups...)
		return nil
	}, nil
}

// Teams ...
func Teams(teams []string) (Option, error) {
	return func(c *Config) error {
		c.Groups = append(c.Groups, teams...)
		return nil
	}, nil
}

type Backend interface {
	Keys(ctx context.Context, opts ...Option) ([]string, error)
}

// New ...
func New(ctx context.Context, backend string) (Backend, error) {
	switch backend {
	case "github":
		return NewGitHub(ctx)
	case "stns":
		return NewSTNS()
	default:
		return nil, fmt.Errorf("unsupported backend '%s'", backend)
	}
}

func unique(in []string) []string {
	m := map[string]struct{}{}
	for _, s := range in {
		m[s] = struct{}{}
	}
	u := []string{}
	for _, s := range in { // fixed order
		if _, ok := m[s]; !ok {
			continue
		}
		u = append(u, s)
	}
	return u
}
