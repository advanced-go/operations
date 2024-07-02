package landscape1

import "time"

// All changes need to go through this database data. Changes can be pulled by different agents
// Types of changes:
// 1. Update - main entry status, assignee class, assigned region.

// EntryChange - add an agentID?
type EntryChange struct {
	Partition int       `json:"partition"`
	AgentId   string    `json:"agency-id"`
	CreatedTS time.Time `json:"created-ts"`

	Item     string `json:"item"` // "status", "class", "region"
	NewValue string `json:"new-value"`
}
