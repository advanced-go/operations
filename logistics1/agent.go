package logistics1

import (
	"fmt"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

const (
	Class = "logistics1"
)

type logistics struct {
	running             bool
	uri                 string
	region              string
	interval            time.Duration
	caseOfficerInterval time.Duration
	ctrlC               chan *messaging.Message
	caseOfficers        *messaging.Exchange
	shutdown            func()
}

func AgentUri(region string) string {
	return fmt.Sprintf("%v:%v", Class, region)
}

// NewAgent - create a new logistics agent, region needs to be set via host environment
func NewAgent(region string) messaging.Agent {
	return newAgent(region)
}

// newAgent - create a new logistics agent
func newAgent(region string) *logistics {
	c := new(logistics)
	c.uri = AgentUri(region)
	c.region = region
	c.interval = time.Second * 2
	c.caseOfficerInterval = time.Second * 5
	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.caseOfficers = messaging.NewExchange()
	return c
}

// String - identity
func (l *logistics) String() string {
	return l.uri
}

// Uri - agent identifier
func (l *logistics) Uri() string {
	return l.uri
}

// Message - message the agent
func (l *logistics) Message(m *messaging.Message) {
	messaging.Mux(m, l.ctrlC, nil, nil)
}

// Shutdown - shutdown the agent
func (l *logistics) Shutdown() {
	if !l.running {
		return
	}
	l.running = false
	if l.shutdown != nil {
		l.shutdown()
	}
	msg := messaging.NewControlMessage(l.uri, l.uri, messaging.ShutdownEvent)
	if l.ctrlC != nil {
		l.ctrlC <- msg
	}
	l.caseOfficers.Broadcast(msg)
}

// Run - run the agent
func (l *logistics) Run() {
	if l.running {
		return
	}
	l.running = true

	go run(l, activity1.Log, queryAssignments, newCaseOfficerAgent)
}
