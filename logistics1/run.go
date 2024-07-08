package logistics1

import (
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

type logFunc func(body []activity1.Entry) *core.Status
type agentFunc func(interval time.Duration, traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent
type getFunc func(region string) ([]landscape1.Entry, *core.Status)

// run - operations logistics
func run(l *logistics, log logFunc, get getFunc, agent agentFunc) {
	if l == nil {
		return
	}
	init := false
	tick := time.Tick(l.interval)

	for {
		select {
		case <-tick:
			// TODO : determine how to check for partition changes
		case msg, open := <-l.ctrlC:
			if !open {
				return
			}
			switch msg.Event() {
			case messaging.ShutdownEvent:
				close(l.ctrlC)
				return
			default:
			}
		default:
			if !init {
				init = true
				status := processAssignments(l, log, get, agent)
				if !status.OK() && !status.NotFound() {
					log([]activity1.Entry{{AgentId: l.uri, Details: status.Err.Error()}})
					// TODO : how to handle log error
				}
			}
		}
	}
}
