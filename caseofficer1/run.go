package caseofficer1

import (
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/url"
	"time"
)

// run - case officer run function
func run(c *caseOfficer, update func(partition landscape1.Entry) ([]assignment1.Entry, *core.Status)) {
	if c == nil {
		return
	}
	if update == nil {
		update = updateAssignments
	}
	tick := time.Tick(c.interval)
	for {
		select {
		case <-tick:
			status := c.log([]activity1.Entry{{AgentId: c.agentId}})
			if !status.OK() {
				c.parent.Message(messaging.NewMessageWithStatus(messaging.ChannelData, "to", "from", "event", status))
			} else {
				entries, status1 := update(c.partition)
				if !status1.OK() {
					c.parent.Message(messaging.NewMessageWithStatus(messaging.ChannelData, "to", "from", "event", status1))
				} else {
					// Assign new ingress and egress agents
					if status1.OK() && len(entries) > 0 {
					}
				}
			}
		case msg, open := <-c.agentCh:
			if !open {
				return
			}
			switch msg.Event() {
			case messaging.ShutdownEvent:
				return
			default:
			}
		default:
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
