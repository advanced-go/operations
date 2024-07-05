package caseofficer1

import (
	"fmt"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

const (
	Class = "case-officer1"
)

type caseOfficer struct {
	running       bool
	uri           string
	interval      time.Duration
	partition     landscape1.Entry
	ctrlC         chan *messaging.Message
	statusCtrlC   chan *messaging.Message
	statusC       chan *messaging.Message
	handler       messaging.Agent
	ingressAgents *messaging.Exchange
	egressAgents  *messaging.Exchange
	shutdown      func()
}

func AgentUri(traffic string, origin core.Origin) string {
	if origin.SubZone == "" {
		return fmt.Sprintf("%v:%v.%v.%v", Class, traffic, origin.Region, origin.Zone)
	}
	return fmt.Sprintf("%v:%v.%v.%v.%v", Class, traffic, origin.Region, origin.Zone, origin.SubZone)
}

func AgentUriFromAssignment(e landscape1.Entry) string {
	return AgentUri(e.Traffic, e.Origin())
}

// NewAgent - create a new case officer agent
func NewAgent(interval time.Duration, partition landscape1.Entry, handler messaging.Agent) messaging.Agent {
	return newAgent(interval, partition, handler)
}

// newAgent - create a new case officer agent
func newAgent(interval time.Duration, partition landscape1.Entry, handler messaging.Agent) *caseOfficer {
	c := new(caseOfficer)
	c.uri = AgentUriFromAssignment(partition)
	c.interval = interval
	c.partition = partition
	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.statusCtrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.statusC = make(chan *messaging.Message, 3*messaging.ChannelSize)
	c.handler = handler
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
	c.shutdown = messaging.AddShutdown(c.shutdown, f)
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
	c.running = true
	go runStatus(c, logActivity, insertAssignmentStatus)
	go run(c, logActivity, updateAssignments, newControllerAgent)
}
