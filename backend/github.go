package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GitHub struct {
	client *githubv4.Client
}

// NewGitHub ...
func NewGitHub(ctx context.Context) (*GitHub, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(ctx, src)

	endpoint := os.Getenv("GITHUB_ENDPOINT")
	if endpoint == "" {
		endpoint = "https://api.github.com/graphql"
	}

	log.Printf("backend endpoint: %s\n", endpoint)

	return &GitHub{
		client: githubv4.NewEnterpriseClient(endpoint, httpClient),
	}, nil
}

func (b *GitHub) Keys(ctx context.Context, options ...Option) ([]string, error) {
	users := []string{}
	keys := []string{}
	c := &Config{}
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	for _, u := range c.Users {
		var q struct {
			User struct {
				PublicKeys struct {
					Edges []struct {
						Node struct {
							Key githubv4.String
						}
					}
				} `graphql:"publicKeys(first: 100)"`
			} `graphql:"user(login: $user)"`
		}
		variables := map[string]interface{}{
			"user": githubv4.String(u),
		}
		users = append(users, u)

		if err := b.client.Query(ctx, &q, variables); err != nil {
			return nil, err
		}
		for _, k := range q.User.PublicKeys.Edges {
			keys = append(keys, string(k.Node.Key))
		}
	}

	// org teams
	for _, g := range c.Groups {
		splited := strings.Split(g, "/")
		if len(splited) != 2 {
			return nil, fmt.Errorf("invalid GitHub org team name '%s'", g)
		}
		org := splited[0]
		team := splited[1]
		var q struct {
			Org struct {
				Team struct {
					Members struct {
						Nodes []struct {
							Login      githubv4.String
							PublicKeys struct {
								Edges []struct {
									Node struct {
										Key githubv4.String
									}
								}
							} `graphql:"publicKeys(first: 100)"`
						}
					}
				} `graphql:"team(slug: $team)"`
			} `graphql:"organization(login: $org)"`
		}
		variables := map[string]interface{}{
			"org":  githubv4.String(org),
			"team": githubv4.String(team),
		}

		if err := b.client.Query(ctx, &q, variables); err != nil {
			return nil, err
		}
		for _, u := range q.Org.Team.Members.Nodes {
			users = append(users, string(u.Login))
			for _, k := range u.PublicKeys.Edges {
				keys = append(keys, string(k.Node.Key))
			}
		}
	}

	log.Printf("public keys collected from GitHub %s\n", unique(users))

	return unique(keys), nil
}
