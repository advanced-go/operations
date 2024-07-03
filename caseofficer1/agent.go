package caseofficer1

import (
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

const (
	class = "case-officer"
)

type caseOfficer struct {
	running       bool
	uri           string
	agentId       string
	interval      time.Duration
	partition     landscape1.Entry
	agentCh       chan *messaging.Message
	statusCh      chan *messaging.Message
	parent        messaging.Agent
	ingressAgents *messaging.Exchange
	egressAgents  *messaging.Exchange
	shutdown      func()
	log           func(body []activity1.Entry) *core.Status
}

// NewAgent - create a new case officer agent
func NewAgent(uri string, interval time.Duration, partition landscape1.Entry, parent messaging.Agent) messaging.Agent {
	return newAgent(uri, interval, partition, parent, func(body []activity1.Entry) *core.Status {
		_, status := activity1.Put(nil, body)
		return status
	})
}

// newAgent - create a new case officer agent
func newAgent(uri string, interval time.Duration, partition landscape1.Entry, parent messaging.Agent, log func(body []activity1.Entry) *core.Status) messaging.Agent {
	c := new(caseOfficer)
	c.uri = uri
	c.agentId = class + ":" + uri
	c.interval = interval
	c.partition = partition

	c.agentCh = make(chan *messaging.Message, messaging.ChannelSize)
	c.statusCh = make(chan *messaging.Message, messaging.ChannelSize)
	c.parent = parent
	c.ingressAgents = messaging.NewExchange()
	c.egressAgents = messaging.NewExchange()

	c.log = log
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
func (c *caseOfficer) Message(msg *messaging.Message) {
	if msg == nil {
		return
	}
	if msg.Channel() == messaging.ChannelControl && c.agentCh != nil {
		c.agentCh <- msg
	}
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
	c.Message(messaging.NewControlMessage(c.uri, c.uri, messaging.ShutdownEvent))
}

// Run - run the agent
func (c *caseOfficer) Run() {
	if c.running {
		return
	}
	//TODO : start status processing
	//TODO : read all existing assignments and create agents
	//TODO : start run function
}
