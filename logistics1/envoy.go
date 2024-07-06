package logistics1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
)

const (
	Class = "envoy1"
)

type envoy struct {
	running bool
	uri     string
	//interval      time.Duration
	//partition     landscape1.Entry
	ctrlC chan *messaging.Message
	//statusCtrlC   chan *messaging.Message
	//statusC       chan *messaging.Message
	//handler       messaging.Agent
	caseOfficers *messaging.Exchange
	//egressAgents  *messaging.Exchange
	shutdown func()
}

func AgentUri(traffic string, origin core.Origin) string {
	if origin.SubZone == "" {
		return fmt.Sprintf("%v:%v.%v.%v", Class, traffic, origin.Region, origin.Zone)
	}
	return fmt.Sprintf("%v:%v.%v.%v.%v", Class, traffic, origin.Region, origin.Zone, origin.SubZone)
}

//func AgentUriFromAssignment(e landscape1.Entry) string {
//	return AgentUri(e.Traffic, e.Origin())
//}

// NewEnvoyAgent - create a new envoy agent
func NewEnvoyAgent() messaging.Agent {
	return newEnvoyAgent()
}

// newEnvoyAgent - create a new envoy agent
func newEnvoyAgent() *envoy {
	c := new(envoy)
	c.uri = "" //AgentUriFromAssignment(partition)
	//c.interval = interval
	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	//c.statusCtrlC = make(chan *messaging.Message, messaging.ChannelSize)
	//c.statusC = make(chan *messaging.Message, 3*messaging.ChannelSize)
	//c.handler = handler
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
