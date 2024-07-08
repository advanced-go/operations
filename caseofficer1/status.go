package caseofficer1

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
)

type insertFunc func(msg *messaging.Message) *core.Status

// run - case officer status processing
func runStatus(c *caseOfficer, log logFunc, insert insertFunc) {
	if c == nil {
		return
	}
	for {
		select {
		case msg, open := <-c.statusC:
			if !open {
				return
			}
			status1 := log(nil, c.uri, "processing status message")
			if !status1.OK() {
				c.handler.Message(messaging.NewStatusMessage(c.handler.Uri(), c.uri, status1))
			} else {
				status1 = insert(msg)
				c.handler.Message(messaging.NewStatusMessage(c.handler.Uri(), c.uri, status1))
				if !status1.OK() && !status1.NotFound() {
					c.handler.Message(messaging.NewStatusMessage(c.handler.Uri(), c.uri, status1))
				}
			}
		case msg1, open1 := <-c.statusCtrlC:
			if !open1 {
				return
			}
			switch msg1.Event() {
			case messaging.ShutdownEvent:
				log(nil, c.uri, "shutting down")
				close(c.statusC)
				close(c.statusCtrlC)
				return
			default:
			}
		default:
		}
	}
}
