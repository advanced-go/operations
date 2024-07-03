package caseofficer1

import (
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/url"
	"time"
)

// run - case officer run function
func run(c *caseOfficer, fnTick func(officer *caseOfficer) ([]assignment1.Entry, *core.Status)) {
	if c == nil {
		return
	}
	if fnTick == nil {
		fnTick = onTick
	}
	tick := time.Tick(c.interval)
	for {
		select {
		case <-tick:
			status := c.log([]activity1.Entry{{AgentId: c.agentId}})
			if !status.OK() {
				c.parent.Message(messaging.NewMessageWithStatus(messaging.ChannelData, "to", "from", "event", status))
			}
			entries, status1 := fnTick(c)
			if !status1.OK() {
				c.parent.Message(messaging.NewMessageWithStatus(messaging.ChannelData, "to", "from", "event", status1))
			}
			// Assign new ingress and egress agents
			if len(entries) > 0 {
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

func onTick(c *caseOfficer) ([]assignment1.Entry, *core.Status) {
	values := make(url.Values)
	values.Add(core.RegionKey, c.partition.Region)
	values.Add(core.ZoneKey, c.partition.Zone)
	values.Add(core.SubZoneKey, c.partition.SubZone)
	values.Add("traffic", c.partition.Traffic)
	entries, _, status := assignment1.Get(nil, nil, values)
	return entries, status
}
