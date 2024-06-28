package landscape1

import "time"

var (
	entryData = []Entry{
		{Partition: 1, Region: "us-west1", Zone: "a", SubZone: "", Traffic: "ingress", Status: "active", AssigneeClass: "ingress-case-class1", AssignedRegion: "west", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 2, Region: "us-west1", Zone: "b", SubZone: "", Traffic: "egress", Status: "active", AssigneeClass: "egress-case-class1", AssignedRegion: "west", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 3, Region: "us-south1", Zone: "c", SubZone: "", Traffic: "ingress", Status: "active", AssigneeClass: "ingress-case-class1", AssignedRegion: "west", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 4, Region: "us-south1", Zone: "d", SubZone: "", Traffic: "egress", Status: "active", AssigneeClass: "egress-case-class1", AssignedRegion: "east", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 5, Region: "us-central1", Zone: "a", SubZone: "", Traffic: "ingress", Status: "active", AssigneeClass: "ingress-case-class1", AssignedRegion: "east", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 6, Region: "us-central1", Zone: "c", SubZone: "", Traffic: "egress", Status: "active", AssigneeClass: "egress-case-class1", AssignedRegion: "east", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Entry - agent
type Entry struct {
	Partition int       `json:"partition"`
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Traffic   string    `json:"traffic"`
	Status    string    `json:"status"` // active or inactive
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	AssigneeClass  string `json:"assignee-class"`
	AssignedRegion string `json:"assigned-region"` // "east" or "west"

}

// EntryChange - add an agentID?
type EntryChange struct {
	Partition int       `json:"partition"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	Item     string `json:"item"` // "status", "class", "region"
	NewValue string `json:"new-value"`
}

// HostEntry - listing of all hosts
type HostEntry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Assignment
	Host string `json:"host"`
	Type string `json:"type"` // backbone vs operations

}
