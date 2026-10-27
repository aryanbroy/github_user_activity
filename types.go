package main

type Repo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PR struct {
	Type string `json:"type"`
	Repo Repo   `json:"repo"`
}

type Event string

const (
	CreateEvent      Event = "CreateEvent"
	PushEvent        Event = "PushEvent"
	IssuesEvent      Event = "IssuesEvent"
	WatchEvent       Event = "WatchEvent"
	PullRequestEvent Event = "PullRequestEvent"
)
