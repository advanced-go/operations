package logistics1

import (
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/caseofficer1"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/url"
	"time"
)

type logFunc func(body []activity1.Entry) *core.Status
type agentFunc func(interval time.Duration, traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent
type getFunc func(region string) ([]landscape1.Entry, *core.Status)

// run - operations envoy
func run(e *envoy, log logFunc, get getFunc, agent agentFunc) {
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
				status := processAssignments(e, log, get, agent)
				if !status.OK() && !status.NotFound() {
					log([]activity1.Entry{{AgentId: e.uri, Details: status.Err.Error()}})
					// TODO : how to handle log error
				}
			}
		}
	}
}

func logActivity(body []activity1.Entry) *core.Status {
	_, status := activity1.Put(nil, body)
	return status
}

func getAssignments(region string) ([]landscape1.Entry, *core.Status) {
	values := make(url.Values)
	values.Add(landscape1.AssignedRegionKey, region)
	values.Add(landscape1.StatusKey, landscape1.StatusActive)
	return landscape1.Get(nil, nil, values)
}

func processAssignments(e *envoy, log logFunc, get getFunc, newAgent agentFunc) *core.Status {
	status := log([]activity1.Entry{{AgentId: e.uri}})
	if !status.OK() {
		return status
	}
	entries, status1 := get(e.region)
	if !status1.OK() {
		return status1
	}
	for _, e1 := range entries {
		e.caseOfficers.Register(newAgent(e.caseOfficerInterval, e1.Traffic, e1.Origin(), e))
	}
	return status
}

func newAgent(interval time.Duration, traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent {
	return caseofficer1.NewAgent(interval, traffic, origin, handler)
}
