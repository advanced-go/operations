package logistics1

import (
	"fmt"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/caseofficer1"
	"github.com/advanced-go/operations/landscape1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/url"
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
	c.interval = time.Second * 5
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

// Add - add a shutdown function
//func (l *logistics) Add(f func()) {
//	l.shutdown = messaging.AddShutdown(l.shutdown, f)
//}

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
func (e *logistics) Run() {
	if e.running {
		return
	}
	e.running = true

	go run(e, logActivity, getAssignments, newCaseOfficerAgent)
}

func logActivity(body []activity1.Entry) *core.Status {
	_, status := activity1.Put(nil, body)
	return status
}

func getAssignments(region string) ([]landscape1.Entry, *core.Status) {
	values := make(url.Values)
	values.Add(landscape1.AssignedRegionKey, region)
	values.Add(landscape1.StatusKey, landscape1.StatusActive)
	return landscape1.Get(nil, nil, values)
}

func newCaseOfficerAgent(interval time.Duration, traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent {
	return caseofficer1.NewAgent(interval, traffic, origin, handler)
}

func processAssignments(l *logistics, log logFunc, get getFunc, newAgent agentFunc) *core.Status {
	status := log([]activity1.Entry{{AgentId: l.uri}})
	if !status.OK() {
		return status
	}
	entries, status1 := get(l.region)
	if !status1.OK() {
		return status1
	}
	for _, e1 := range entries {
		l.caseOfficers.Register(newAgent(l.caseOfficerInterval, e1.Traffic, e1.Origin(), l))
	}
	return status
}
