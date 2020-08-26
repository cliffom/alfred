package main

import (
	"fmt"
	"log"
	"time"

	"github.com/icrowley/fake"
	"github.com/spf13/viper"
)

type repos struct {
	Tag   string
	Env   string
	Repos map[string]repo
}

type repo struct {
	Owner       string
	Branch      string
	Prerelease  bool
	Draft       bool
	Title       string
	Description string
}

func getRepos(cfg *viper.Viper) *repos {
	r := &repos{}
	if err := cfg.Unmarshal(&r); err != nil {
		log.Fatal(fmt.Errorf("Could not process config file: %s", err))
	}

	if !(len(r.Tag) > 0) {
		r.Tag = generateTag(r.Env)
	}
	return r
}

func getReposConfig(env string) *viper.Viper {
	reposConfig := viper.New()
	reposConfig.SetConfigName(env)
	reposConfig.AddConfigPath(".alfred")

	if err := reposConfig.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("%s", err))
	}

	reposConfig.Set("Env", env)
	return reposConfig
}

func generateTag(env string) string {
	t := time.Now()
	tagBase := string(t.Format("2006.01.02"))
	return fmt.Sprintf("%s-%s.%s", env, tagBase, fake.Word())
}
