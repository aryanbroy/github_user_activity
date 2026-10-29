package main

import (
	"fmt"
	"strings"
)

func HandleCreateEvent(repoName string) string {
	resp := fmt.Sprintf("Created a new branch in %s", repoName)
	return resp
}

func HandleWatchEvent(repoName string) string {
	resp := fmt.Sprintf("Starred %s", repoName)
	return resp
}
func HandleIssuesEvent() string {

	var resp string

	for repoInf, count := range issueMap {
		res := strings.Split(repoInf, "/")
		repoName := res[0] + "/" + res[1]
		resAction := res[2]
		if count > 1 {
			switch resAction {
			case "opened":
				resp = fmt.Sprintf("Opened %v issues in %s", count, repoName)
			case "closed":
				resp = fmt.Sprintf("Closed %v issues in %s", count, repoName)
			case "reopened":
				resp = fmt.Sprintf("Reopened %v issues in %s", count, repoName)
			case "assigned":
				resp = fmt.Sprintf("Assigned %v issues in %s", count, repoName)
			case "unassigned":
				resp = fmt.Sprintf("Unassigned %v issues in %s", count, repoName)
			default:
				resp = fmt.Sprintf("Action not implemented yet...")
			}
		} else {
			switch resAction {
			case "opened":
				resp = fmt.Sprintf("Opened an issue in %s", repoName)
			case "closed":
				resp = fmt.Sprintf("Closed an issue in %s", repoName)
			case "reopened":
				resp = fmt.Sprintf("Reopened an issue in %s", repoName)
			case "assigned":
				resp = fmt.Sprintf("Assigned an issue in %s", repoName)
			case "unassigned":
				resp = fmt.Sprintf("Unassigned an issue in %s", repoName)
			case "edited":
				resp = fmt.Sprintf("Edited an issue in %s", repoName)
			case "deleted":
				resp = fmt.Sprintf("Deleted an issue in %s", repoName)
			case "labeled":
				resp = fmt.Sprintf("Added a label to an issue in %s", repoName)
			case "unlabeled":
				resp = fmt.Sprintf("Removed a label from an issue in %s", repoName)
			case "locked":
				resp = fmt.Sprintf("Locked an issue in %s", repoName)
			case "unlocked":
				resp = fmt.Sprintf("Unlocked an issue in %s", repoName)
			case "milestoned":
				resp = fmt.Sprintf("Added an issue to a milestone in %s", repoName)
			case "demilestoned":
				resp = fmt.Sprintf("Removed an issue from a milestone in %s", repoName)
			case "transferred":
			default:
				resp = fmt.Sprintf("Action not implemented yet...")
			}
		}
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
