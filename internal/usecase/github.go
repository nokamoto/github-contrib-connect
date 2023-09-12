package usecase

import (
	"context"
	"time"

	"github.com/nokamoto/github-contrib-connect/internal/infra"
	"github.com/nokamoto/github-contrib-connect/pkg/apis"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Contribs struct {
	github *infra.Github
}

func NewContribs(github *infra.Github) *Contribs {
	return &Contribs{
		github: github,
	}
}

func (c *Contribs) ListContribs(ctx context.Context, org string, repos []string, since time.Time) (*apis.Contrib, error) {
	var contrib apis.Contrib
	for _, repo := range repos {
		pulls, err := c.github.ListPulls(ctx, org, repo, since)
		if err != nil {
			return nil, err
		}

		for _, pull := range pulls {
			num := pull.GetNumber()
			comments, err := c.github.ListComments(ctx, org, repo, num)
			if err != nil {
				return nil, err
			}

			repository := &apis.Repository{
				Org:  org,
				Name: repo,
			}

			for _, comment := range comments {
				contrib.Comments = append(contrib.GetComments(), &apis.Comment{
					Repository: repository,
					User:       comment.GetUser().GetLogin(),
					CreatedAt:  timestamppb.New(comment.GetCreatedAt().Time),
				})
			}

			reviews, err := c.github.ListReviews(ctx, org, repo, num)
			if err != nil {
				return nil, err
			}

			for _, review := range reviews {
				contrib.Reviews = append(contrib.GetReviews(), &apis.Review{
					Repository: repository,
					User:       review.GetUser().GetLogin(),
					CreatedAt:  timestamppb.New(review.GetSubmittedAt().Time),
					State:      review.GetState(),
				})
			}
		}
	}
	return &contrib, nil
}
