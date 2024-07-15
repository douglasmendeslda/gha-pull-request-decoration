package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	// Get environment variables
	repoOwner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	repoName := os.Getenv("GITHUB_REPOSITORY_NAME")
	prNumber := os.Getenv("GITHUB_PR_NUMBER")
	githubToken := os.Getenv("GITHUB_TOKEN")

	if repoOwner == "" || repoName == "" || prNumber == "" || githubToken == "" {
		fmt.Println("Some environment variables are missing")
		return
	}

	// Create a GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Create a comment
	comment := &github.IssueComment{
		Body: github.String("Hello World"),
	}
	_, _, err := client.Issues.CreateComment(ctx, repoOwner, repoName, prNumber, comment)
	if err != nil {
		fmt.Printf("Error creating comment: %v\n", err)
		return
	}

	fmt.Println("Comment created successfully")
}
