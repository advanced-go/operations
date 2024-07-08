package caseofficer1

import (
	"context"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

type logFunc func(body []activity1.Entry) *core.Status
type updateFunc func(ctx context.Context, agentId string, origin core.Origin) ([]assignment1.Entry, *core.Status)
type agentFunc func(traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent

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
