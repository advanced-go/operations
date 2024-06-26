package landscape1

import "time"

// Entry - host
type Entry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Assignment
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`

	// Region + Zone + Class
	AssigneeTag string `json:"assignee-tag"` // Assigned to an agent class and origin

}

// EntryStatus - add an agentID?
type EntryStatus struct {
	EntryId      int       `json:"entry-id"`
	AssignmentId int       `json:"assignment-id"`
	AgentId      string    `json:"agent-id"` // Creation agent id
	CreatedTS    time.Time `json:"created-ts"`

	// Status - active/inactive
	East string `json:"east"`
	West string `json:"west"`
}

// Resources - listing of all hosts envoy or operations
type Resources struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Assignment
	Host string `json:"host"`
	Type string `json:"type"` // agency vs operations

}
