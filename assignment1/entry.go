package assignment1

import "time"

// Entry - host
type Entry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Origin
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`
	Host    string `json:"host"`

	// Region + Zone + Class
	AssigneeTag string `json:"assignee-tag"` // Assigned to an agent class and origin

}

// EntryStatus - add an agentID?
type EntryStatus struct {
	EntryId      int       `json:"entry-id"`
	AssignmentId int       `json:"assignment1-id"`
	AgentId      string    `json:"agent-id"` // Creation agent id
	CreatedTS    time.Time `json:"created-ts"`

	// Status - active/inactive
	East string `json:"east"`
	West string `json:"west"`
}
