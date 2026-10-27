package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func ParseResponse(r io.Reader) ([]PR, error) {
	var ps []PR
	err := json.NewDecoder(r).Decode(&ps)
	return ps, err
}

func GithubDataToHumanReadable(rawData PR) {
	eventType := HandleEventsDesc(rawData)
	if eventType == "error" {
		log.Fatalln("Bad event")
	}

	repoName := rawData.Repo.Name
	fmt.Println("Repo name: ", repoName)
	fmt.Println("Event type: ", eventType)
}

func HandleEventsDesc(gitData PR) string {
	event := Event(gitData.Type)
	switch event {
	case CreateEvent:
		return "create"
	case PushEvent:
		return "push"
	case IssuesEvent:
		return "issue"
	case WatchEvent:
		return "watch"
	case PullRequestEvent:
		return "pullReq"
	default:
		fmt.Println("Invalid event or the project is small enough to not record this event")
		return "error"
	}
}
