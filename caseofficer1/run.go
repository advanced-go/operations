package caseofficer1

import (
	"context"
	"github.com/advanced-go/agency/egress1"
	"github.com/advanced-go/agency/ingress1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

type logFunc func(ctx context.Context, agentId string, content any) *core.Status
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
			status := processAssignments(c, update, agent)
			log(nil, c.uri, "process assignments : onTick()")
			if !status.OK() && !status.NotFound() {
				c.handler.Message(messaging.NewStatusMessage(c.handler.Uri(), c.uri, status))
			}
		case msg, open := <-c.ctrlC:
			if !open {
				return
			}
			switch msg.Event() {
			case messaging.ShutdownEvent:
				log(nil, c.uri, messaging.ShutdownEvent)
				close(c.ctrlC)
				return
			default:
			}
		default:
			if !init {
				init = true
				status := processAssignments(c, update, agent)
				log(nil, c.uri, "process assignments : default")
				if !status.OK() && !status.NotFound() {
					c.handler.Message(messaging.NewStatusMessage(c.handler.Uri(), c.uri, status))
				}
				log(nil, c.uri, "process assignments : default")
			}
		}
	}
}

func newControllerAgent(traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent {
	if traffic == access.IngressTraffic {
		return ingress1.NewControllerAgent(origin, handler)
	}
	return egress1.NewControllerAgent(origin, handler)
}

func processAssignments(c *caseOfficer, update updateFunc, newAgent agentFunc) *core.Status {
	//status := log(nil, c.uri, "processingAssignment")
	//if !status.OK() {
	//	return status
	//}
	entries, status := update(nil, c.uri, c.origin)
	if !status.OK() {
		return status
	}
	for _, e := range entries {
		c.controllers.Register(newAgent(c.traffic, e.Origin(), c.handler))
	}
	return status
}
