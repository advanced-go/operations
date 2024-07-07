package caseofficer1

import (
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
)

type insertFunc func(msg *messaging.Message) *core.Status

// run - status processing
func runStatus(c *caseOfficer, log logFunc, insert insertFunc) {
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
				close(c.statusC)
				close(c.statusCtrlC)
				return
			default:
			}
		default:
		}
	}
}
