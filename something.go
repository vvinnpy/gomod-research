package gomod_research

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
)

func PrintVersion() {
	fmt.Println(`v4.1.0, del go.mod and import with v4.1.0+incompatible`)
}

func FetchOrganizations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

func FetchWatched(username string) ([]*github.Repository, *github.Response, error) {
	client := github.NewClient(nil)
	watched, r, err := client.Activity.ListWatched(context.Background(), username, nil)
	if err != nil {
		return nil, nil, err
	}
	return watched, r, err
}
