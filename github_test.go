package gomod_research

import (
	"fmt"
	"testing"
)

func TestOrganizations(t *testing.T) {
	var username = "vvinnpy"
	fmt.Print("Enter GitHub username: ")
	organizations, err := FetchOrganizations(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, organization := range organizations {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}
}

func TestRepositories(t *testing.T) {
	var username = "vvinnpy"
	fmt.Print("Enter GitHub username: ")
	watched, _, err := FetchWatched(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for i, w := range watched {
		fmt.Printf("%v. %v\n", i+1, w)
	}
}
