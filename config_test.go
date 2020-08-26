package main

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	githubToken := fake.CharactersN(16)
	notifyURL := fake.DomainName()
	cfg := &config{}

	cfgViper := viper.New()
	cfgViper.Set("GithubToken", githubToken)
	cfgViper.Set("NotifyURL", notifyURL)
	cfgViper.Unmarshal(&cfg)

	if cfg.GithubToken != githubToken {
		t.Errorf("Value of GithubToken was incorrect, got: %s, expected: %s", cfg.GithubToken, githubToken)
	}

	if cfg.NotifyURL != notifyURL {
		t.Errorf("Value of GithubToken was incorrect, got: %s, expected: %s", cfg.GithubToken, githubToken)
	}
}
