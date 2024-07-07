package caseofficer1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/agency/egress1"
	"github.com/advanced-go/agency/ingress1"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
	"net/url"
	"time"
)

const (
	Class = "case-officer1"
)

type caseOfficer struct {
	running     bool
	uri         string
	interval    time.Duration
	traffic     string
	origin      core.Origin
	ctrlC       chan *messaging.Message
	statusCtrlC chan *messaging.Message
	statusC     chan *messaging.Message
	handler     messaging.Agent
	controllers *messaging.Exchange
	shutdown    func()
}

func AgentUri(traffic string, origin core.Origin) string {
	if origin.SubZone == "" {
		return fmt.Sprintf("%v:%v.%v.%v", Class, traffic, origin.Region, origin.Zone)
	}
	return fmt.Sprintf("%v:%v.%v.%v.%v", Class, traffic, origin.Region, origin.Zone, origin.SubZone)
}

// NewAgent - create a new case officer agent
func NewAgent(interval time.Duration, traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent {
	return newCaseAgent(interval, traffic, origin, handler)
}

// newCaseAgent - create a new case officer agent
func newCaseAgent(interval time.Duration, traffic string, origin core.Origin, handler messaging.Agent) *caseOfficer {
	c := new(caseOfficer)
	c.uri = AgentUri(traffic, origin)
	c.traffic = traffic
	c.origin = origin
	c.interval = interval

	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.statusCtrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.statusC = make(chan *messaging.Message, 3*messaging.ChannelSize)
	c.handler = handler
	c.controllers = messaging.NewExchange()
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
	c.controllers.Broadcast(msg)
}

// Run - run the agent
func (c *caseOfficer) Run() {
	if c.running {
		return
	}
	c.running = true
	go runStatus(c, logActivity, insertAssignmentStatus)
	go run(c, logActivity, updateAssignments, newAgent)
}

func logActivity(body []activity1.Entry) *core.Status {
	req, _ := http.NewRequest("", "https://www.google.com/search?q=golang", nil)
	_, status := activity1.Put(req, body)
	return status
}

func insertAssignmentStatus(msg *messaging.Message) *core.Status {
	status := msg.Status()
	if status == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("message body content is not of type *core.Status"))
	}
	values := make(url.Values)
	values.Add(core.RegionKey, msg.Header.Get(core.RegionKey))
	values.Add(core.ZoneKey, msg.Header.Get(core.ZoneKey))
	values.Add(core.SubZoneKey, msg.Header.Get(core.SubZoneKey))
	values.Add(core.HostKey, msg.Header.Get(core.HostKey))
	return assignment1.InsertStatus(nil, values, status)
}

func updateAssignments(origin core.Origin) ([]assignment1.Entry, *core.Status) {
	values := make(url.Values)
	values.Add(core.RegionKey, origin.Region)
	values.Add(core.ZoneKey, origin.Zone)
	values.Add(core.SubZoneKey, origin.SubZone)
	//values.Add(core.HostKey, origin.Host)
	entries, _, status := assignment1.Get(nil, nil, values)
	return entries, status
}

func newAgent(traffic string, origin core.Origin, handler messaging.Agent) messaging.Agent {
	if traffic == access.IngressTraffic {
		return ingress1.NewControllerAgent(origin, handler)
	}
	return egress1.NewControllerAgent(origin, handler)
}

func processAssignments(c *caseOfficer, log logFunc, update updateFunc, newAgent agentFunc) *core.Status {
	status := log([]activity1.Entry{{AgentId: c.uri}})
	if !status.OK() {
		return status
	}
	entries, status1 := update(c.origin)
	if !status1.OK() {
		return status
	}
	for _, e := range entries {
		c.controllers.Register(newAgent(c.traffic, e.Origin(), c.handler))
	}
	return status
}
