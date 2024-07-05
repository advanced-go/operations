package caseofficer1

import (
	"errors"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/url"
)

// run - status processing
func runStatus(c *caseOfficer, log func(body []activity1.Entry) *core.Status, insert func(msg *messaging.Message) *core.Status) {
	if c == nil {
		return
	}
	for {
		select {
		case msg, open1 := <-c.statusC:
			if !open1 {
				return
			}
			status1 := log([]activity1.Entry{{AgentId: c.uri}})
			if !status1.OK() {
				c.handler.Message(messaging.NewStatusMessage("", "", "", status1))
			} else {
				status1 = insert(msg)
				if !status1.OK() && !status1.NotFound() {
					c.handler.Message(messaging.NewStatusMessage("", "", "", status1))
				}
			}
		case msg, open := <-c.ctrlC:
			if !open {
				return
			}
			switch msg.Event() {
			case messaging.ShutdownEvent:
				return
			default:
			}
		default:
		}
	}
}

func insertAssignmentStatus(msg *messaging.Message) *core.Status {
	status := msg.Status()
	if status == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("message body content is not of type *core.Status"))
	}
	values := make(url.Values)
	values.Add(core.RegionKey, msg.Header.Get(core.RegionKey))
	values.Add(core.ZoneKey, msg.Header.Get(core.ZoneKey))
	values.Add(core.SubZoneKey, msg.Header.Get(core.SubZoneKey))
	values.Add(core.HostKey, msg.Header.Get(core.HostKey))
	return assignment1.InsertStatus(nil, values, status)
}
