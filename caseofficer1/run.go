package caseofficer1

import (
	"github.com/advanced-go/agency/egress1"
	"github.com/advanced-go/agency/ingress1"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/url"
	"time"
)

type logFunc func(body []activity1.Entry) *core.Status
type updateFunc func(traffic string, origin core.Origin) ([]assignment1.Entry, *core.Status)
type agentFunc func(traffic string, entry assignment1.Entry, handler messaging.Agent) messaging.Agent

// run - case officer
func run(c *caseOfficer, log logFunc, update updateFunc, agent agentFunc) {
	if c == nil {
		return
	}
	init := false
	tick := time.Tick(c.interval)
	for {
		select {
		case <-tick:
			status := processAssignments(c, log, update, agent)
			if !status.OK() && !status.NotFound() {
				c.handler.Message(messaging.NewStatusMessage("", "", "", status))
			}
		case msg, open := <-c.ctrlC:
			if !open {
				return
			}
			switch msg.Event() {
			case messaging.ShutdownEvent:
				close(c.ctrlC)
				return
			default:
			}
		default:
			if !init {
				init = true
				status := processAssignments(c, log, update, agent)
				if !status.OK() && !status.NotFound() {
					c.handler.Message(messaging.NewStatusMessage("", "", "", status))
				}
			}
		}
	}
}

func logActivity(body []activity1.Entry) *core.Status {
	_, status := activity1.Put(nil, body)
	return status
}

func updateAssignments(traffic string, origin core.Origin) ([]assignment1.Entry, *core.Status) {
	values := make(url.Values)
	values.Add("traffic", traffic)
	values.Add(core.RegionKey, origin.Region)
	values.Add(core.ZoneKey, origin.Zone)
	values.Add(core.SubZoneKey, origin.SubZone)
	entries, _, status := assignment1.Get(nil, nil, values)
	return entries, status
}

func processAssignments(c *caseOfficer, log logFunc, update updateFunc, newAgent agentFunc) *core.Status {
	status := log([]activity1.Entry{{AgentId: c.uri}})
	if !status.OK() {
		return status
	}
	entries, status1 := update(c.traffic, c.origin)
	if !status1.OK() {
		return status
	}
	for _, e := range entries {
		c.controllers.Register(newAgent(c.traffic, e, c.handler))
	}
	return status
}

func newAgent(traffic string, entry assignment1.Entry, handler messaging.Agent) messaging.Agent {
	if traffic == access.IngressTraffic {
		return ingress1.NewAgent(entry.Origin(), handler)
	}
	return egress1.NewAgent(entry.Origin(), handler)
}
