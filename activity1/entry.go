package activity1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/operations/common"
	"net/url"
	"time"
)

const (
	EntryIdName   = "entry_id"
	AgentIdName   = "agent_id"
	CreatedTSName = "created_ts"
	ContentName   = "content"
)

var (
	safeEntry = common.NewSafe()
	entryData = []Entry{
		{EntryId: 1, AgentId: "agency-id", Content: "testing 1-2-3", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Entry - agency
type Entry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agency-id"`
	CreatedTS time.Time `json:"created-ts"`

	Content string `json:"content"`
}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)
		case AgentIdName:
			e.AgentId = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case ContentName:
			e.Content = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Entry) Values() []any {
	return []any{
		e.EntryId,
		e.AgentId,
		e.CreatedTS,
		e.Content,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func validEntry(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	// Additional filtering
	return true
}

func lastEntry() Entry {
	return entryData[len(entryData)-1]
}
