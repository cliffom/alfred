package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	GithubToken string
	NotifyURL   string
	Environment string
}

func getConfig(env string) *config {
	cfg := viper.New()

	cfg.SetConfigName("config")
	cfg.AddConfigPath(".alfred")

	if err := cfg.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %s", err))
	}

	c := &config{}
	cfg.Unmarshal(&c)
	c.Environment = env

	return c
}
