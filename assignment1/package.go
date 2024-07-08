package assignment1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
	"time"
)

const (
	PkgPath            = "github/advanced-go/operations/assignment1"
	assignmentResource = "assignment"
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	return get[core.Log, Entry](ctx, h, values, assignmentResource, "", nil)
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (h2 http.Header, status *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status1 := json2.New[[]Entry](r.Body, r.Header)
		if !status1.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status1
		}
		body = content
	}
	switch p := any(&body).(type) {
	case *[]Entry:
		h2, status = put[core.Log, Entry](r.Context(), core.AddRequestId(r.Header), assignmentResource, "", *p, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(body))
	}
	return
}

// InsertStatus - insert an assignment status
func InsertStatus(ctx context.Context, agentId string, origin core.Origin, status *core.Status) *core.Status {
	s := EntryStatus{
		Region:    origin.Region,
		Zone:      origin.Zone,
		SubZone:   origin.SubZone,
		Host:      origin.Host,
		AgentId:   agentId,
		CreatedTS: time.Now().UTC(),
		Code:      status.Code,
		Content:   "",
	}
	if status.Err != nil {
		s.Content = status.Err.Error()
	} else {
		if status.Content != nil {
			s.Content = fmt.Sprintf("%v", status.Content)
		} else {
			s.Content = status.String()
		}
	}
	safeStatus.Lock()()
	statusData = append(statusData, s)
	return core.StatusOK()
}

// Update - update the assignments based on new access host entries, returning the new assignments
// TODO: Need to distinguish startup from processing as startup needs the entire list of assignments.
//
//	Normal processing just needs the changes
func Update(ctx context.Context, agentId string, origin core.Origin) ([]Entry, *core.Status) {
	values := make(url.Values)
	values.Add(core.RegionKey, origin.Region)
	values.Add(core.ZoneKey, origin.Zone)
	values.Add(core.SubZoneKey, origin.SubZone)
	entries, _, status := Get(nil, nil, values)
	return entries, status
}
