package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"
	"time"

	"github.com/google/go-github/v55/github"
	"github.com/nokamoto/github-contrib-connect/internal/infra"
	"github.com/nokamoto/github-contrib-connect/pkg/apis"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Repository struct {
	github *infra.Github
	logger *slog.Logger
}

type PullFilter func(*github.PullRequest) string

func match(f []PullFilter, pull *github.PullRequest) []string {
	var matches []string
	for _, filter := range f {
		if why := filter(pull); why != "" {
			matches = append(matches, why)
		}
	}
	return matches
}

func AuthorPullFilter(re string) PullFilter {
	return func(pull *github.PullRequest) string {
		if regexp.MustCompile(re).MatchString(pull.GetUser().GetLogin()) {
			return fmt.Sprintf("author %s matches %s", pull.GetUser().GetLogin(), re)
		}
		return ""
	}
}

func NewRepository(github *infra.Github, logger *slog.Logger) *Repository {
	return &Repository{
		github: github,
		logger: logger,
	}
}

func (r *Repository) List(ctx context.Context, org string, repos []string, since time.Time, filters ...PullFilter) ([]*apis.Repository, error) {
	var repositories []*apis.Repository
	for _, repo := range repos {
		logger := r.logger.With(slog.String("org", org), slog.String("repo", repo))

		repository := &apis.Repository{
			Org:  org,
			Name: repo,
		}

		pulls, err := r.github.ListPulls(ctx, org, repo, since)
		if err != nil {
			return nil, err
		}

		logger.Debug("list pulls", slog.Int("pulls", len(pulls)))

		for _, pull := range pulls {
			logger = logger.With(
				slog.Int("number", pull.GetNumber()),
				slog.String("author", pull.GetUser().GetLogin()),
				slog.String("title", pull.GetTitle()),
			)

			if why := match(filters, pull); len(why) != 0 {
				logger.Debug("skip pull request", slog.Any("why", why))
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

			logger.Debug("list reviews", slog.Int("reviews", len(reviews)))

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
