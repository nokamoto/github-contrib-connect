package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v55/github"
)

type Github struct {
	client *github.Client
	sleep  time.Duration
}

func NewClient(token string, sleep time.Duration) *Github {
	return &Github{
		client: github.NewClient(nil).WithAuthToken(token),
		sleep:  sleep,
	}
}

func (g *Github) ListRepositories(ctx context.Context, org string) ([]*github.Repository, error) {
	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var repositories []*github.Repository
	for {
		repos, res, err := g.client.Repositories.ListByOrg(ctx, org, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list repositories: %s: %w", org, err)
		}
		repositories = append(repositories, repos...)
		if res.NextPage == 0 {
			break
		}
		opts.Page = res.NextPage
		time.Sleep(g.sleep)
	}
	return repositories, nil
}

func (g *Github) ListPulls(ctx context.Context, owner, repo string, since time.Time) ([]*github.PullRequest, error) {
	opts := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var pulls []*github.PullRequest
	for {
		ps, res, err := g.client.PullRequests.List(ctx, owner, repo, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list pull requests: %s/%s: %w", owner, repo, err)
		}
		eop := false
		for _, p := range ps {
			if p.GetCreatedAt().Before(since) {
				eop = true
				break
			}
			pulls = append(pulls, p)
		}
		if eop || res.NextPage == 0 {
			break
		}
		opts.Page = res.NextPage
		time.Sleep(g.sleep)
	}
	return pulls, nil
}

func (g *Github) ListReviews(ctx context.Context, owner, repo string, number int) ([]*github.PullRequestReview, error) {
	opts := &github.ListOptions{PerPage: 100}
	var reviews []*github.PullRequestReview
	for {
		rs, res, err := g.client.PullRequests.ListReviews(ctx, owner, repo, number, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list pull request reviews: %s/%s#%d: %w", owner, repo, number, err)
		}
		reviews = append(reviews, rs...)
		if res.NextPage == 0 {
			break
		}
		opts.Page = res.NextPage
		time.Sleep(g.sleep)
	}
	return reviews, nil
}

func (g *Github) ListComments(ctx context.Context, owner, repo string, number int) ([]*github.PullRequestComment, error) {
	opts := &github.PullRequestListCommentsOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var comments []*github.PullRequestComment
	for {
		cs, res, err := g.client.PullRequests.ListComments(ctx, owner, repo, number, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list pull request comments: %s/%s#%d: %w", owner, repo, number, err)
		}
		comments = append(comments, cs...)
		if res.NextPage == 0 {
			break
		}
		opts.Page = res.NextPage
		time.Sleep(g.sleep)
	}
	return comments, nil
}
