package caseofficer1

import (
	"github.com/advanced-go/agency/egress1"
	"github.com/advanced-go/agency/ingress1"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/url"
	"time"
)

// run - case officer
func run(c *caseOfficer, log func(body []activity1.Entry) *core.Status, update func(partition landscape1.Entry) ([]assignment1.Entry, *core.Status), agent func(traffic string, entry assignment1.Entry, parent messaging.Agent) messaging.Agent) {
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

func updateAssignments(partition landscape1.Entry) ([]assignment1.Entry, *core.Status) {
	values := make(url.Values)
	values.Add(core.RegionKey, partition.Region)
	values.Add(core.ZoneKey, partition.Zone)
	values.Add(core.SubZoneKey, partition.SubZone)
	entries, _, status := assignment1.Get(nil, nil, values)
	return entries, status
}

func logActivity(body []activity1.Entry) *core.Status {
	_, status := activity1.Put(nil, body)
	return status
}

func processAssignments(c *caseOfficer, log func(body []activity1.Entry) *core.Status, update func(partition landscape1.Entry) ([]assignment1.Entry, *core.Status), newAgent func(traffic string, entry assignment1.Entry, parent messaging.Agent) messaging.Agent) *core.Status {
	status := log([]activity1.Entry{{AgentId: c.uri}})
	if !status.OK() {
		return status
	}
	entries, status1 := update(c.partition)
	if !status1.OK() {
		return status
	}
	for _, e := range entries {
		c.ingressAgents.Register(newAgent(access.IngressTraffic, e, c.handler))
		c.egressAgents.Register(newAgent(access.EgressTraffic, e, c.handler))
	}
	return status
}

func newControllerAgent(traffic string, entry assignment1.Entry, parent messaging.Agent) messaging.Agent {
	if traffic == access.IngressTraffic {
		return ingress1.NewAgent(entry.Origin(), parent)
	}
	return egress1.NewAgent(entry.Origin(), parent)
}
