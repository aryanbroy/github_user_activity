package main

import (
	"fmt"
)

func HandleCreateEvent(repoName string) {
	fmt.Printf("Created a new branch in %s\n", repoName)
}

func HandleWatchEvent(repoName string) {
	fmt.Printf("Starred %s\n", repoName)
}
func HandleIssuesEvent(action Action, repoName string, count int) {

	if count > 1 {
		switch action {
		case Opened:
			fmt.Printf("Opened %v issues in %s\n", count, repoName)
		case Closed:
			fmt.Printf("Closed %v issues in %s\n", count, repoName)
		case Reopened:
			fmt.Printf("Reopened %v issues in %s\n", count, repoName)
		case Assigned:
			fmt.Printf("Assigned %v issues in %s\n", count, repoName)
		case Unassigned:
			fmt.Printf("Unassigned %v issues in %s\n", count, repoName)
		default:
			fmt.Println("Action not implemented yet...")
		}
	} else {
		switch action {
		case Opened:
			fmt.Printf("Opened an issue in %s\n", repoName)
		case Closed:
			fmt.Printf("Closed an issue in %s\n", repoName)
		case Reopened:
			fmt.Printf("Reopened an issue in %s\n", repoName)
		case Assigned:
			fmt.Printf("Assigned an issue in %s\n", repoName)
		case Unassigned:
			fmt.Printf("Unassigned an issue in %s\n", repoName)
		case Edited:
			fmt.Printf("Edited an issue in %s\n", repoName)
		case Deleted:
			fmt.Printf("Deleted an issue in %s\n", repoName)
		case Labeled:
			fmt.Printf("Added a label to an issue in %s\n", repoName)
		case Unlabeled:
			fmt.Printf("Removed a label from an issue in %s\n", repoName)
		case Locked:
			fmt.Printf("Locked an issue in %s\n", repoName)
		case Unlocked:
			fmt.Printf("Unlocked an issue in %s\n", repoName)
		case Milestoned:
			fmt.Printf("Added an issue to a milestone in %s\n", repoName)
		case Demilestoned:
			fmt.Printf("Removed an issue from a milestone in %s\n", repoName)
		case Transferred:
			fmt.Printf("Transferred an issue from %s\n", repoName)
		default:
			fmt.Println("Action not implemented yet...")
		}
	}
}

func HandlePullEvent(repoName string) {
	fmt.Printf("Opened a pull request in %s\n", repoName)
}

func HandlePushEvent(repoName string, count int) {
	if count == 1 {
		fmt.Printf("Pushed a commit to %s\n", repoName)
	} else if count > 1 {
		fmt.Printf("Pushed %v commits to %s\n", count, repoName)
	} else {
		fmt.Printf("Count value is invalid, count: %v\n,", count)
	}
}
