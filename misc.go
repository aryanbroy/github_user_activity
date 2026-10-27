package main

import "fmt"

func HandleWatchEvent(repoName string) string {
	desc := fmt.Sprintf("Starred %s", repoName)
	return desc
}
