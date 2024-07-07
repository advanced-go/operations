package logistics1

import (
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/operations/caseofficer1"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/url"
	"time"
)

type logFunc func(body []activity1.Entry) *core.Status
type agentFunc func(interval time.Duration, traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent

// run - operations envoy
func run(e *envoy, log logFunc, agent agentFunc) {
	if e == nil {
		return
	}
	init := false
	tick := time.Tick(e.interval)

	for {
		select {
		case <-tick:
			// TODO : determine how to check for partition changes
		case msg, open := <-e.ctrlC:
			if !open {
				return
			}
			switch msg.Event() {
			case messaging.ShutdownEvent:
				close(e.ctrlC)
				return
			default:
			}
		default:
			if !init {
				init = true
				processPartitions(e, log, agent)
				//if !status.OK() && !status.NotFound() {
				//	   c.handler.Message(messaging.NewStatusMessage("", "", "", status))
				//  }
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

func processPartitions(c *envoy, log logFunc, newAgent agentFunc) *core.Status {
	status := log([]activity1.Entry{{AgentId: c.uri}})
	if !status.OK() {
		return status
	}
	//entries, status1 := update(c.traffic, c.origin)
	//if !status1.OK() {
	//		return status/
	//	}
	//	for _, e := range entries {
	//		c.controllerAgents.Register(newAgent(c.traffic, e, c.handler))
	//	}
	return status
}

func newAgent(interval time.Duration, traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent {
	return caseofficer1.NewAgent(interval, traffic, origin, handler)
}
