package logistics1

import (
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

const (
	Class = "envoy1"
)

type envoy struct {
	running             bool
	uri                 string
	region              string
	interval            time.Duration
	caseOfficerInterval time.Duration
	ctrlC               chan *messaging.Message
	caseOfficers        *messaging.Exchange
	shutdown            func()
}

// NewEnvoyAgent - create a new envoy agent
func NewEnvoyAgent() messaging.Agent {
	return newEnvoyAgent()
}

// newEnvoyAgent - create a new envoy agent
func newEnvoyAgent() *envoy {
	c := new(envoy)
	c.uri = Class
	// Needs to be set via host environment
	c.region = "west"
	c.interval = time.Second * 5
	c.caseOfficerInterval = time.Second * 5
	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.caseOfficers = messaging.NewExchange()
	return c
}

// String - identity
func (e *envoy) String() string {
	return e.uri
}

// Uri - agent identifier
func (e *envoy) Uri() string {
	return e.uri
}

// Message - message the agent
func (e *envoy) Message(m *messaging.Message) {
	messaging.Mux(m, e.ctrlC, nil, nil)
}

// Add - add a shutdown function
func (e *envoy) Add(f func()) {
	e.shutdown = messaging.AddShutdown(e.shutdown, f)
}

// Shutdown - shutdown the agent
func (e *envoy) Shutdown() {
	if !e.running {
		return
	}
	e.running = false
	if e.shutdown != nil {
		e.shutdown()
	}
	msg := messaging.NewControlMessage(e.uri, e.uri, messaging.ShutdownEvent)
	if e.ctrlC != nil {
		e.ctrlC <- msg
	}
	e.caseOfficers.Broadcast(msg)
}

// Run - run the agent
func (e *envoy) Run() {
	if e.running {
		return
	}
	e.running = true

	go run(e, logActivity, getAssignments, newAgent)
}
