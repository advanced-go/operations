package logistics1

import (
	"context"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

type logFunc func(ctx context.Context, agentId string, content any) *core.Status
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
					log(nil, l.uri, status.Err)
					// TODO : how to handle log error
				}
			}
		}
	}
}
