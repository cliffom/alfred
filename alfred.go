package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	env := flag.String("env", "dev", "Environment to create release for")
	testMode := flag.Bool("test-mode", true, "Test mode")
	flag.Parse()

	config := getConfig(*env)
	reposConfig := getReposConfig(config.Environment)
	repos := getRepos(reposConfig)

	ctx := context.Background()
	client := getClient(ctx, config.GithubToken)

	fmt.Println("Test Mode: ", *testMode)
	for name, repo := range repos.Repos {
		rel := newRelease(repos.Tag, repo)
		var message *message
		if !*testMode {
			rel, _, err := client.Repositories.CreateRelease(ctx, repo.Owner, name, rel)
			if err != nil {
				log.Fatal("Could not create release: ", err)
			}
			message = newMessage(name, config.Environment, rel)
			message.notify(config.NotifyURL)
		} else {
			message = newMessage(name, config.Environment, rel)
			fmt.Println(message.pretty())
		}
	}
}

func newMessage(repo string, env string, rel *github.RepositoryRelease) *message {
	return &message{
		Repo:        repo,
		Environment: env,
		Release:     rel,
	}
}

func newRelease(t string, r repo) *github.RepositoryRelease {
	author := "alfred"
	return &github.RepositoryRelease{
		TagName:         &t,
		TargetCommitish: &r.Branch,
		Name:            &r.Title,
		Body:            &r.Description,
		Draft:           &r.Draft,
		Prerelease:      &r.Prerelease,
		Author: &github.User{
			Login: &author,
		},
	}
}

func getClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}
