package main

import "fmt"

func HandleCreateEvent(repoName string) string {
	resp := fmt.Sprintf("Created a new branch in %s", repoName)
	return resp
}

func HandleWatchEvent(repoName string) string {
	resp := fmt.Sprintf("Starred %s", repoName)
	return resp
}

func HandleIssuesEvent(repoName string, payload Payload) string {
	action := payload.Action

	var resp string

	switch action {
	case Opened:
		resp = fmt.Sprintf("Opened an issue in %s", repoName)
	case Closed:
		resp = fmt.Sprintf("Closed an issue in %s", repoName)
	case Reopened:
		resp = fmt.Sprintf("Reopened an issue in %s", repoName)
	case Assigned:
		resp = fmt.Sprintf("Assigned an issue in %s", repoName)
	case Unassigned:
		resp = fmt.Sprintf("Unassigned an issue in %s", repoName)
	case Edited:
		resp = fmt.Sprintf("Edited an issue in %s", repoName)
	case Deleted:
		resp = fmt.Sprintf("Deleted an issue in %s", repoName)
	case Labeled:
		resp = fmt.Sprintf("Added a label to an issue in %s", repoName)
	case Unlabeled:
		resp = fmt.Sprintf("Removed a label from an issue in %s", repoName)
	case Locked:
		resp = fmt.Sprintf("Locked an issue in %s", repoName)
	case Unlocked:
		resp = fmt.Sprintf("Unlocked an issue in %s", repoName)
	case Milestoned:
		resp = fmt.Sprintf("Added an issue to a milestone in %s", repoName)
	case Demilestoned:
		resp = fmt.Sprintf("Removed an issue from a milestone in %s", repoName)
	case Transferred:
		resp = fmt.Sprintf("Transferred an issue from %s", repoName)
	default:
		resp = fmt.Sprintf("Action not implemented yet...")
	}

	return resp
}

func HandlePullEvent(repoName string) string {
	resp := fmt.Sprintf("Opened a pull request in %s", repoName)
	return resp
}

func HandlePushEvent(repoName string) string {
	resp := fmt.Sprintf("Pushed a commit to %s", repoName)
	return resp
}
