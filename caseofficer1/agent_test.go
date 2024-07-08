package caseofficer1

import (
	"fmt"
	"github.com/advanced-go/operations/activity1"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
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

func ExampleInsertAssignmentStatus() {
	msg := messaging.NewMessageWithStatus(messaging.ChannelStatus, "to", "from", "", core.StatusOK())
	status := insertAssignmentStatus(msg)

	fmt.Printf("test: insertAssignmentStatus() -> [status:%v]\n", status)

	//Output:
	//test: insertAssignmentStatus() -> [status:OK]

}

func ExampleNewControllerAgent() {
	origin := core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "",
		Host:       "www.host1.com",
		InstanceId: "",
	}
	a := newControllerAgent(access.IngressTraffic, origin, nil)
	fmt.Printf("test: newControllerAgent(\"%v\") -> [%v]\n", access.IngressTraffic, a)

	a = newControllerAgent(access.EgressTraffic, origin, nil)
	fmt.Printf("test: newControllerAgent(\"%v\") -> [%v]\n", access.EgressTraffic, a)

	//Output:
	//test: newControllerAgent("ingress") -> [ingress-controller1:us-central1.c.www.host1.com]
	//test: newControllerAgent("egress") -> [egress-controller1:us-central1.c.www.host1.com]

}

func ExampleProcessAssignments() {
	origin := core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "",
		Host:       "www.host1.com",
		InstanceId: "",
	}

	c := newAgent(time.Second*5, access.IngressTraffic, origin, nil)
	fmt.Printf("test: newAgent() -> [status:%v]\n", c != nil)

	status := processAssignments(c, activity1.Log, assignment1.Update, newControllerAgent)
	fmt.Printf("test: processAssignments() -> [status:%v] [controllers:%v]\n", status, c.controllers.Count())

	//Output:
	//test: newAgent() -> [status:true]
	//test: processAssignments() -> [status:OK] [controllers:2]

}
