package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func ParseResponse(r io.Reader) ([]PR, error) {
	var ps []PR
	err := json.NewDecoder(r).Decode(&ps)
	return ps, err
}

func GithubDataToHumanReadable(rawData PR) {
	desc, typ := HandleEventsDesc(rawData)
	if typ == "issue" {
		fmt.Println(desc)
	}

	// fmt.Println(desc)
	// repoName := rawData.Repo.Name
	// fmt.Println("Repo name: ", repoName)
	// fmt.Println("Event type: ", eventType)
}

func HandleEventsDesc(gitData PR) (string, string) {
	event := Event(gitData.Type)
	repoName := gitData.Repo.Name
	payload := gitData.Payload

	switch event {
	case CreateEvent:
		return HandleCreateEvent(repoName), "create"
	case PushEvent:
		return HandlePushEvent(repoName), "push"
	case IssuesEvent:
		return HandleIssuesEvent(repoName, payload), "issue"
	case WatchEvent:
		return HandleWatchEvent(repoName), "watch"
	case PullRequestEvent:
		return HandlePullEvent(repoName), "pull"
	default:
		resp := fmt.Sprintf("Invalid event or the project is small enough to not record this event: ", gitData.Type)
		return resp, "error"
	}
}
