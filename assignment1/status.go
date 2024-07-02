package assignment1

import (
	"errors"
	"fmt"
	"time"
)

const (

	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	statusSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"

	statusInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst,retry,retry_rate_limit,retry_rate_burst,failover) VALUES"

	StatusName = "status"
)

var (
	statusData = []EntryStatus{
		{Region: "us-west-2", Zone: "usw2-az4", Host: "www.host2.com", Status: "error", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-west-2", Zone: "usw2-az3", Host: "www.host1.com", Status: "other", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

func lastStatus() EntryStatus {
	return statusData[len(statusData)-1]
}

// EntryStatus - add an agentID?
type EntryStatus struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	CreatedTS time.Time `json:"created-ts"`

	// Status information for non-standard assignment processing, such as error. Case Officer should monitor this
	Status string `json:"status"`
}

func (EntryStatus) Scan(columnNames []string, values []any) (e EntryStatus, err error) {
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
		case StatusName:
			e.Status = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e EntryStatus) Values() []any {
	return []any{
		e.Region,
		e.Zone,
		e.SubZone,
		e.Host,
		e.CreatedTS,
		e.Status,
	}
}

func (EntryStatus) Rows(entries []EntryStatus) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
