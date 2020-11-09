package backend

import (
	"context"
	"fmt"
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
	var client *githubv4.Client
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(ctx, src)

	switch {
	case os.Getenv("GITHUB_ENDPOINT") != "":
		client = githubv4.NewEnterpriseClient(os.Getenv("GITHUB_ENDPOINT"), httpClient)
	default:
		client = githubv4.NewClient(httpClient)
	}

	return &GitHub{
		client: client,
	}, nil
}

func (b *GitHub) Keys(ctx context.Context, options ...Option) ([]string, error) {
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
			for _, k := range u.PublicKeys.Edges {
				keys = append(keys, string(k.Node.Key))
			}
		}
	}

	return unique(keys), nil
}
