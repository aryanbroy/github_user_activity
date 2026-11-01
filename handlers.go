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
func HandleIssuesEvent() {

	for repoInf, count := range issueMap {
		res := strings.Split(repoInf, "/")
		repoName := res[0] + "/" + res[1]
		resAction := res[2]
		if count > 1 {
			switch resAction {
			case "opened":
				fmt.Printf("Opened %v issues in %s\n", count, repoName)
			case "closed":
				fmt.Printf("Closed %v issues in %s\n", count, repoName)
			case "reopened":
				fmt.Printf("Reopened %v issues in %s\n", count, repoName)
			case "assigned":
				fmt.Printf("Assigned %v issues in %s\n", count, repoName)
			case "unassigned":
				fmt.Printf("Unassigned %v issues in %s\n", count, repoName)
			default:
				fmt.Println("Action not implemented yet...")
			}
		} else {
			switch resAction {
			case "opened":
				fmt.Printf("Opened an issue in %s\n", repoName)
			case "closed":
				fmt.Printf("Closed an issue in %s\n", repoName)
			case "reopened":
				fmt.Printf("Reopened an issue in %s\n", repoName)
			case "assigned":
				fmt.Printf("Assigned an issue in %s\n", repoName)
			case "unassigned":
				fmt.Printf("Unassigned an issue in %s\n", repoName)
			case "edited":
				fmt.Printf("Edited an issue in %s\n", repoName)
			case "deleted":
				fmt.Printf("Deleted an issue in %s\n", repoName)
			case "labeled":
				fmt.Printf("Added a label to an issue in %s\n", repoName)
			case "unlabeled":
				fmt.Printf("Removed a label from an issue in %s\n", repoName)
			case "locked":
				fmt.Printf("Locked an issue in %s\n", repoName)
			case "unlocked":
				fmt.Printf("Unlocked an issue in %s\n", repoName)
			case "milestoned":
				fmt.Printf("Added an issue to a milestone in %s\n", repoName)
			case "demilestoned":
				fmt.Printf("Removed an issue from a milestone in %s\n", repoName)
			case "transferred":
				fmt.Printf("Transferred an issue from %s\n", repoName)
			default:
				fmt.Println("Action not implemented yet...")
			}
		}
	}
}

func HandlePullEvent(repoName string) string {
	resp := fmt.Sprintf("Opened a pull request in %s", repoName)
	return resp
}

func HandlePushEvent(repoName string) string {
	resp := fmt.Sprintf("Pushed a commit to %s", repoName)
	return resp
}
