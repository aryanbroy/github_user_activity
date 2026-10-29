package main

type Repo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PR struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Action Action `json:"action"`
}

type Event string

const (
	CreateEvent      Event = "CreateEvent"
	PushEvent        Event = "PushEvent"
	IssuesEvent      Event = "IssuesEvent"
	WatchEvent       Event = "WatchEvent"
	PullRequestEvent Event = "PullRequestEvent"
)

type Action string

const (
	Opened       Action = "opened"
	Edited       Action = "edited"
	Deleted      Action = "deleted"
	Closed       Action = "closed"
	Reopened     Action = "reopened"
	Assigned     Action = "assigned"
	Unassigned   Action = "unassigned"
	Labeled      Action = "labeled"
	Unlabeled    Action = "unlabeled"
	Locked       Action = "locked"
	Unlocked     Action = "unlocked"
	Milestoned   Action = "milestoned"
	Demilestoned Action = "demilestoned"
	Transferred  Action = "transferred"
)
