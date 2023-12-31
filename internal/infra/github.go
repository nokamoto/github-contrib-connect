package infra

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/go-github/v55/github"
)

type Github struct {
	client *github.Client
	sleep  time.Duration
	logger *slog.Logger
}

func NewClient(token string, sleep time.Duration, logger *slog.Logger) *Github {
	return &Github{
		client: github.NewClient(nil).WithAuthToken(token),
		sleep:  sleep,
		logger: logger,
	}
}

func NewEnterpriseClient(baseURL, token string, sleep time.Duration, logger *slog.Logger) (*Github, error) {
	e, err := github.NewClient(nil).WithEnterpriseURLs(baseURL, baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create a github client: %s: %w", baseURL, err)
	}
	return &Github{
		client: e.WithAuthToken(token),
		sleep:  sleep,
		logger: logger,
	}, nil
}

func (g *Github) ListRepositories(ctx context.Context, org string) ([]*github.Repository, error) {
	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var repositories []*github.Repository
	for {
		time.Sleep(g.sleep)
		repos, res, err := g.client.Repositories.ListByOrg(ctx, org, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list repositories: %s: %w", org, err)
		}
		repositories = append(repositories, repos...)
		if res.NextPage == 0 {
			break
		}
		opts.Page = res.NextPage
	}
	return repositories, nil
}

func (g *Github) ListPulls(ctx context.Context, owner, repo string, since time.Time) ([]*github.PullRequest, error) {
	logger := g.logger.With(slog.String("owner", owner), slog.String("repo", repo), slog.String("since", since.String()))

	opts := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var pulls []*github.PullRequest
	for {
		time.Sleep(g.sleep)
		logger.Debug("list pull requests", slog.Int("page", opts.Page))
		ps, res, err := g.client.PullRequests.List(ctx, owner, repo, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list pull requests: %s/%s: %w", owner, repo, err)
		}
		eop := false
		for _, p := range ps {
			if p.GetCreatedAt().Before(since) {
				logger.Debug("end of pull requests", slog.Int("number", p.GetNumber()), slog.String("created_at", p.GetCreatedAt().String()))
				eop = true
				break
			}
			pulls = append(pulls, p)
		}
		if eop || res.NextPage == 0 {
			break
		}
		opts.Page = res.NextPage
	}
	return pulls, nil
}

func (g *Github) ListReviews(ctx context.Context, owner, repo string, number int) ([]*github.PullRequestReview, error) {
	logger := g.logger.With(slog.String("owner", owner), slog.String("repo", repo), slog.Int("number", number))

	opts := &github.ListOptions{PerPage: 100}
	var reviews []*github.PullRequestReview
	for {
		time.Sleep(g.sleep)
		logger.Debug("list pull request reviews", slog.Int("page", opts.Page))
		rs, res, err := g.client.PullRequests.ListReviews(ctx, owner, repo, number, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list pull request reviews: %s/%s#%d: %w", owner, repo, number, err)
		}
		reviews = append(reviews, rs...)
		if res.NextPage == 0 {
			break
		}
		opts.Page = res.NextPage
	}
	return reviews, nil
}
