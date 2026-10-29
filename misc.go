package main

var issueMap = make(map[string]int)

func CountIssues(gitData []PR) {

	for _, data := range gitData {
		if data.Type == "IssuesEvent" {
			repoName := data.Repo.Name
			action := data.Payload.Action

			issueMap[repoName+"/"+string(action)] += 1
		}
	}
}
