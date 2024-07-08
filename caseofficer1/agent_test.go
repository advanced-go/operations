package caseofficer1

import (
	"fmt"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
)

func ExampleAgentUri() {
	u := AgentUri("ingress", core.Origin{Region: "us-central1", Zone: "c", SubZone: "sub-zone"})
	fmt.Printf("test: AgentUri() -> [%v]\n", u)

	u = AgentUri("egress", core.Origin{Region: "us-west1", Zone: "a"})
	fmt.Printf("test: AgentUri() -> [%v]\n", u)

	//Output:
	//test: AgentUri() -> [case-officer1:ingress.us-central1.c.sub-zone]
	//test: AgentUri() -> [case-officer1:egress.us-west1.a]

}

func ExampleNewAgent() {
	// a := NewAgent()
	fmt.Printf("test: newAgent() -> ")

	//Output:
	//test: newAgent() ->

}

func ExampleLogActivity() {
	status := activity1.Log(nil, "agent-id", "example log ")

	fmt.Printf("test: activity1.Log() -> [status:%v]\n", status)

	//Output:
	//test: activity1.Log() -> [status:OK]

}

type testAgent struct{}

func newTestAgent() *testAgent {
	return new(testAgent)
}

func (t *testAgent) Uri() string { return "testAgent" }

func (t *testAgent) Message(m *messaging.Message) {
	//if m.Channel() == messaging.ChannelStatus {
	//	status := m.Status()
	//	fmt.Printf("test: testAgent.Message() -> [status:%v] %v\n", status, m)
	//} else {
	fmt.Printf("test: testAgent.Message() -> %v\n", m)

}

func (t *testAgent) Run() {}

func (t *testAgent) Shutdown() {}
