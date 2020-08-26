package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

type message struct {
	Repo        string
	Environment string
	Release     *github.RepositoryRelease
}

func (m *message) oneline() string {
	return fmt.Sprintf(
		"[%s] Release %s@%s (%s) created by %s.",
		m.Environment,
		m.Repo,
		*m.Release.TagName,
		*m.Release.TargetCommitish,
		*m.Release.Author.Login)
}

func (m *message) pretty() string {
	message := fmt.Sprintf(
		"Release `%s@%s (%s)` created by `%s`\n\n%s\n\n%s\n=-=-=-=-=-=-=-=-=-=-\n\n",
		m.Repo,
		*m.Release.TagName,
		*m.Release.TargetCommitish,
		*m.Release.Author.Login,
		*m.Release.Name,
		*m.Release.Body)

	return message
}

func (m *message) notify(url string) {
	log.Printf(m.oneline())
	if !*m.Release.Prerelease {
		jsonStr := fmt.Sprintf("{\"text\": \"%s\"}", m.pretty())
		data := []byte(jsonStr)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, _ := client.Do(req)
		resp.Body.Close()
	}
}
