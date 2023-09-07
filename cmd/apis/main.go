package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nokamoto/github-contrib-connect/internal/infra"
)

func main() {
	client := infra.NewClient(os.Getenv("GITHUB_TOKEN"), 1*time.Second)
	ctx := context.Background()

	repos, err := client.ListRepositories(ctx, "golang")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1; i++ {
		pulls, err := client.ListPulls(ctx, "golang", repos[i].GetName())
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s/%s: %d\n", repos[i].GetOwner().GetLogin(), repos[i].GetName(), len(pulls))
	}
}
