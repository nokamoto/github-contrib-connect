package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/nokamoto/github-contrib-connect/internal/infra"
)

type Contribs struct {
	github *infra.Github
}

func NewContribs(github *infra.Github) *Contribs {
	return &Contribs{
		github: github,
	}
}

func (c *Contribs) ListContribs(ctx context.Context, org string, repos []string, since time.Time) error {
	for _, repo := range repos {
		pulls, err := c.github.ListPulls(ctx, org, repo, since)
		if err != nil {
			return err
		}

		for _, pull := range pulls {
			num := pull.GetNumber()
			comments, err := c.github.ListComments(ctx, org, repo, num)
			if err != nil {
				return err
			}

			for _, comment := range comments {
			}

			reviews, err := c.github.ListReviews(ctx, org, repo, num)
			if err != nil {
				return err
			}

			for _, review := range reviews {
				review.GetState()
			}
		}
	}
}
