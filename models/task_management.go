package models

import "time"

type TaskItems struct {
	Items             []Task      `json:"items"`
	ContinuationToken interface{} `json:"continuationToken"`
}
type Task struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Message       string    `json:"message"`
	CurrentState  string    `json:"currentState"`
	LastRunResult string    `json:"lastRunResult"`
	NextRun       time.Time `json:"nextRun"`
	LastRun       time.Time `json:"lastRun"`
}

type TaskAttributes struct {
	Action string `json:"action"`
	Method string `json:"method"`
	Data   *Data  `json:"data"`
	Type   string `json:"type"`
	Tid    int    `json:"tid"`
}
type TaskDeleteAttributes struct {
	Action string   `json:"action"`
	Method string   `json:"method"`
	Data   []string `json:"data"`
	Type   string   `json:"type"`
	Tid    int      `json:"tid"`
}
type Properties struct {
	BlobstoreName string `json:"blobstoreName"`
}
type Data []struct {
	ID                    string        `json:"id"`
	TypeID                string        `json:"typeId"`
	Enabled               bool          `json:"enabled"`
	Name                  string        `json:"name"`
	AlertEmail            string        `json:"alertEmail"`
	NotificationCondition string        `json:"notificationCondition"`
	Schedule              string        `json:"schedule"`
	Properties            *Properties   `json:"properties"`
	RecurringDays         []interface{} `json:"recurringDays"`
	StartDate             time.Time     `json:"startDate"`
	TimeZoneOffset        string        `json:"timeZoneOffset"`
}
