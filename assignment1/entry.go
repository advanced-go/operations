package assignment1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/operations/common"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"time"
)

const (
	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	accessLogSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"
	accessLogInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst) VALUES"

	CreatedTSName = "created_ts"
	RegionName    = "region"
	ZoneName      = "zone"
	SubZoneName   = "sub_zone"
	HostName      = "host"
)

var (
	safeEntry = common.NewSafe()
	entryData = []Entry{
		{Region: "us-west1", Zone: "a", Host: "www.host1.com", AgentId: "test-agent", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-west1", Zone: "a", Host: "www.host2.com", AgentId: "test-agent", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host1.com", AgentId: "test-agent", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host2.com", AgentId: "test-agent", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "d", Host: "www.host4.com", AgentId: "test-agent", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

func lastEntry() Entry {
	return entryData[len(entryData)-1]
}

// Case office looks for open assignments, and then does an assignment to a Service Agent

// Entry - host
type Entry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`
}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case RegionName:
			e.Region = values[i].(string)
		case ZoneName:
			e.Zone = values[i].(string)
		case SubZoneName:
			e.SubZone = values[i].(string)
		case HostName:
			e.Host = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Entry) Values() []any {
	return []any{
		e.Region,
		e.Zone,
		e.SubZone,
		e.Host,
		e.CreatedTS,
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
		Host:       e.Host,
		InstanceId: "",
	}
}

func validEntry(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	filter := core.NewOrigin(values)
	target := core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
	if !core.OriginMatch(target, filter) {
		return false
	}
	// Additional filtering
	return true
}
