package landscape1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/operations/common"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"time"
)

const (
	AgentIdName        = "agent_id"
	CreatedTSName      = "created_ts"
	PartitionName      = "partition"
	RegionName         = "region"
	ZoneName           = "zone"
	SubZoneName        = "sub_zone"
	TrafficName        = "traffic"
	StatusName         = "status"
	AssigneeClassName  = "assignee_class"
	AssignedRegionName = "assigned_region"

	StatusKey         = "status"
	StatusActive      = "active"
	TrafficKey        = "traffic"
	AssignedRegionKey = "assigned-region"
)

var (
	safeEntry = common.NewSafe()
	entryData = []Entry{
		{Partition: 1, Region: "us-west1", Zone: "a", SubZone: "", Status: "active", AssigneeClass: "case-officer1", AssignedRegion: "east", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 2, Region: "us-west1", Zone: "b", SubZone: "", Status: "active", AssigneeClass: "case-officer1", AssignedRegion: "east", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 3, Region: "us-south1", Zone: "b", SubZone: "", Status: "active", AssigneeClass: "case-officer1", AssignedRegion: "east", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 4, Region: "us-south1", Zone: "c", SubZone: "", Status: "active", AssigneeClass: "case-officer1", AssignedRegion: "east", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 5, Region: "us-central1", Zone: "c", SubZone: "", Status: "active", AssigneeClass: "case-officer1", AssignedRegion: "west", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 6, Region: "us-central1", Zone: "c", SubZone: "", Status: "active", AssigneeClass: "case-officer1", AssignedRegion: "west", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 7, Region: "us-central1", Zone: "d", SubZone: "", Status: "active", AssigneeClass: "case-officer1", AssignedRegion: "west", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Partition: 8, Region: "us-central1", Zone: "d", SubZone: "", Status: "active", AssigneeClass: "case-officer1", AssignedRegion: "west", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Entry - agency
type Entry struct {
	Partition int       `json:"partition"`
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Status    string    `json:"status"` // active or inactive
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	AssigneeClass  string `json:"assignee-class"`
	AssignedRegion string `json:"assigned-region"` // "east" or "west"

}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case PartitionName:
			e.Partition = values[i].(int)
		case AgentIdName:
			e.AgentId = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)
		case RegionName:
			e.Region = values[i].(string)
		case ZoneName:
			e.Zone = values[i].(string)
		case SubZoneName:
			e.SubZone = values[i].(string)

		//case TrafficName:
		//	e.Traffic = values[i].(string)
		case StatusName:
			e.Status = values[i].(string)

		case AssigneeClassName:
			e.AssigneeClass = values[i].(string)
		case AssignedRegionName:
			e.AssignedRegion = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Entry) Values() []any {
	return []any{
		e.Partition,
		e.AgentId,
		e.CreatedTS,
		e.Region,
		e.Zone,
		e.SubZone,
		//e.Traffic,
		e.Status,

		e.AssigneeClass,
		e.AssignedRegion,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func (e Entry) Origin() core.Origin {
	return core.Origin{
		Region:     e.Region,
		Zone:       e.Zone,
		SubZone:    e.SubZone,
		Host:       "",
		InstanceId: "",
	}
}

func validEntry(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	if isValid(values.Get(StatusKey), e.Status) && isValid(values.Get(AssignedRegionKey), e.AssignedRegion) {
		return true
	}
	return false
}

func isValid(value, target string) bool {
	if value == "" {
		return true
	}
	return value == target
}

func lastEntry() Entry {
	return entryData[len(entryData)-1]
}
