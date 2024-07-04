package caseofficer1

import (
	"fmt"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

const (
	Class = "case-officer"
)

type caseOfficer struct {
	running       bool
	uri           string
	interval      time.Duration
	partition     landscape1.Entry
	ctrlC         chan *messaging.Message
	statusCtrlC   chan *messaging.Message
	statusC       chan *messaging.Message
	parent        messaging.Agent
	ingressAgents *messaging.Exchange
	egressAgents  *messaging.Exchange
	shutdown      func()
}

func AgentUri(traffic, region, zone, subZone string) string {
	if subZone == "" {
		return fmt.Sprintf("%v:%v.%v.%v", Class, traffic, region, zone)
	}
	return fmt.Sprintf("%v:%v.%v.%v.%v", Class, traffic, region, zone, subZone)
}

func AgentUriFromAssignment(entry landscape1.Entry) string {
	if entry.SubZone == "" {
		return fmt.Sprintf("%v:%v.%v.%v", Class, entry.Traffic, entry.Region, entry.Zone)
	}
	return fmt.Sprintf("%v:%v.%v.%v.%v", Class, entry.Traffic, entry.Region, entry.Zone, entry.SubZone)
}

// NewAgent - create a new case officer agent
func NewAgent(uri string, interval time.Duration, partition landscape1.Entry, parent messaging.Agent) messaging.Agent {
	c := new(caseOfficer)
	c.uri = uri
	c.interval = interval
	c.partition = partition

	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.statusCtrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.statusC = make(chan *messaging.Message, messaging.ChannelSize)
	c.parent = parent
	c.ingressAgents = messaging.NewExchange()
	c.egressAgents = messaging.NewExchange()
	return c
}

// String - identity
func (c *caseOfficer) String() string {
	return c.uri
}

// Uri - agent identifier
func (c *caseOfficer) Uri() string {
	return c.uri
}

// Message - message the agent
func (c *caseOfficer) Message(m *messaging.Message) {
	messaging.Mux(m, c.ctrlC, nil, c.statusC)
}

// Add - add a shutdown function
func (c *caseOfficer) Add(f func()) {
	if f == nil {
		return
	}
	if c.shutdown == nil {
		c.shutdown = f
	} else {
		// !panic
		prev := c.shutdown
		c.shutdown = func() {
			prev()
			f()
		}
	}
}

// Shutdown - shutdown the agent
func (c *caseOfficer) Shutdown() {
	if !c.running {
		return
	}
	c.running = false
	if c.shutdown != nil {
		c.shutdown()
	}
	msg := messaging.NewControlMessage(c.uri, c.uri, messaging.ShutdownEvent)
	if c.ctrlC != nil {
		c.ctrlC <- msg
	}
	if c.statusCtrlC != nil {
		c.statusCtrlC <- msg
	}
}

// Run - run the agent
func (c *caseOfficer) Run() {
	if c.running {
		return
	}
	go runStatus(c, c.statusCtrlC, c.statusC, logStatusActivity, insertAssignmentStatus)
	go run(c, c.ctrlC, logActivity, updateAssignments)
}
