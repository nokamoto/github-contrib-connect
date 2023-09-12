package usecase

import (
	"context"
	"regexp"
	"time"

	"github.com/google/go-github/v55/github"
	"github.com/nokamoto/github-contrib-connect/internal/infra"
	"github.com/nokamoto/github-contrib-connect/pkg/apis"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Repository struct {
	github *infra.Github
}

type PullFilter func(*github.PullRequest) bool

func match(f []PullFilter, pull *github.PullRequest) bool {
	for _, filter := range f {
		if filter(pull) {
			return true
		}
	}
	return false
}

func AuthorPullFilter(re string) PullFilter {
	return func(pull *github.PullRequest) bool {
		return regexp.MustCompile(re).MatchString(pull.GetUser().GetLogin())
	}
}

func NewRepository(github *infra.Github) *Repository {
	return &Repository{
		github: github,
	}
}

func (r *Repository) List(ctx context.Context, org string, repos []string, since time.Time, filters ...PullFilter) ([]*apis.Repository, error) {
	var repositories []*apis.Repository
	for _, repo := range repos {
		repository := &apis.Repository{
			Org:  org,
			Name: repo,
		}

		pulls, err := r.github.ListPulls(ctx, org, repo, since)
		if err != nil {
			return nil, err
		}

		for _, pull := range pulls {
			if match(filters, pull) {
				continue
			}

			num := pull.GetNumber()
			pull := &apis.Pull{
				Number:    int32(num),
				User:      pull.GetUser().GetLogin(),
				CreatedAt: timestamppb.New(pull.GetCreatedAt().Time),
			}

			reviews, err := r.github.ListReviews(ctx, org, repo, num)
			if err != nil {
				return nil, err
			}

			for _, review := range reviews {
				pull.Reviews = append(pull.GetReviews(), &apis.Review{
					User:      review.GetUser().GetLogin(),
					CreatedAt: timestamppb.New(review.GetSubmittedAt().Time),
					State:     review.GetState(),
				})
			}

			repository.Pulls = append(repository.GetPulls(), pull)
		}

		repositories = append(repositories, repository)
	}
	return repositories, nil
}
