package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/icrowley/fake"
	"github.com/spf13/viper"
)

func TestReposConfg(t *testing.T) {
	tag := fake.Word()
	owner := fake.FirstName()
	branch := fake.Word()
	prerelease := true
	title := fake.Sentence()
	description := fake.Sentences()

	var repoExpected map[string]repo
	repoExpected = make(map[string]repo)

	repoExpected["repo"] = repo{
		Owner:       owner,
		Branch:      branch,
		Prerelease:  prerelease,
		Title:       title,
		Description: description,
	}
	r := &repos{
		Tag:   tag,
		Repos: repoExpected,
	}
	cfg := &repos{}

	reposConfig := viper.New()
	reposConfig.Set("Tag", tag)
	reposConfig.Set("Repos.repo.Owner", owner)
	reposConfig.Set("Repos.repo.Branch", branch)
	reposConfig.Set("Repos.repo.Prerelease", prerelease)
	reposConfig.Set("Repos.repo.Title", title)
	reposConfig.Set("Repos.repo.Description", description)
	reposConfig.Unmarshal(&cfg)

	if !reflect.DeepEqual(r, cfg) {
		t.Errorf("Configuration does not match expected value.")
	}
}

func TestDefinedTag(t *testing.T) {
	env := fake.Word()
	tag := fake.Word()
	reposConfig := viper.New()

	reposConfig.Set("Tag", tag)
	reposConfig.Set("Env", env)
	repos := getRepos(reposConfig)
	reposConfig.Unmarshal(repos)

	if tag != repos.Tag {
		t.Errorf("Tag does not match expected value, got %s, expected %s",
			repos.Tag, tag)
	}
}

func TestUndefinedTag(t *testing.T) {
	env := fake.Word()
	time := string(time.Now().Format("2006.01.02"))
	tagBase := fmt.Sprintf("%s-%s", env, time)
	reposConfig := viper.New()
	reposConfig.Set("Env", env)
	repos := getRepos(reposConfig)

	if !strings.Contains(repos.Tag, tagBase) {
		t.Errorf("Tag does not match expected value, got %s, expected %s",
			repos.Tag, tagBase)
	}
}
