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
	interval            time.Duration
	caseOfficerInterval time.Duration
	ctrlC               chan *messaging.Message
	//statusCtrlC   chan *messaging.Message
	//statusC       chan *messaging.Message
	//handler       messaging.Agent
	caseOfficers *messaging.Exchange
	//egressAgents  *messaging.Exchange
	shutdown func()
}

// NewEnvoyAgent - create a new envoy agent
func NewEnvoyAgent() messaging.Agent {
	return newEnvoyAgent()
}

// newEnvoyAgent - create a new envoy agent
func newEnvoyAgent() *envoy {
	c := new(envoy)
	c.uri = Class
	c.interval = time.Second * 5
	c.caseOfficerInterval = time.Second * 5
	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.caseOfficers = messaging.NewExchange()
	return c
}

// String - identity
func (c *envoy) String() string {
	return c.uri
}

// Uri - agent identifier
func (c *envoy) Uri() string {
	return c.uri
}

// Message - message the agent
func (c *envoy) Message(m *messaging.Message) {
	messaging.Mux(m, c.ctrlC, nil, nil)
}

// Add - add a shutdown function
func (c *envoy) Add(f func()) {
	c.shutdown = messaging.AddShutdown(c.shutdown, f)
}

// Shutdown - shutdown the agent
func (c *envoy) Shutdown() {
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
	// TODO : need to shutdown case officers
}

// Run - run the agent
func (c *envoy) Run() {
	if c.running {
		return
	}
	c.running = true

	go run(c, logActivity, newAgent)
}
