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

//	func HandleEventsDesc(gitData PR) (string, string) {
//		event := Event(gitData.Type)
//		repoName := gitData.Repo.Name
//		payload := gitData.Payload
//
//		switch event {
//		case CreateEvent:
//			return HandleCreateEvent(repoName), "create"
//		case PushEvent:
//			return HandlePushEvent(repoName), "push"
//		case IssuesEvent:
//			return HandleIssuesEvent(repoName, payload), "issue"
//		case WatchEvent:
//			return HandleWatchEvent(repoName), "watch"
//		case PullRequestEvent:
//			return HandlePullEvent(repoName), "pull"
//		default:
//			resp := fmt.Sprintf("Invalid event or the project is small enough to not record this event: ", gitData.Type)
//			return resp, "error"
//		}
//	}

var eventsMap = make(map[PR]int)

func AggregateEvents(rawData []PR) {
	for _, data := range rawData {
		eventsMap[data]++
	}
}

func GithubDataToHumanReadable() {
	for key, value := range eventsMap {
		HandleEventsDesc(Event(key.Type), key.Payload, key.Repo.Name, value)
	}
}

func HandleEventsDesc(event Event, payload Payload, repoName string, count int) {
	// fmt.Printf("Event: %v, Payload: %v, Reponame: %v, Count: %v\n", event, payload, repoName, count)

	switch event {
	case CreateEvent:
		HandleCreateEvent(repoName)
	case PushEvent:
		HandlePushEvent(repoName)
	case IssuesEvent:
		HandleIssuesEvent(payload.Action, repoName, count)
	case WatchEvent:
		HandleWatchEvent(repoName)
	case PullRequestEvent:
		HandlePullEvent(repoName)
	default:
		fmt.Printf("Invalid event or the project is small enough to not record this event: %s/n", event)
	}
}
