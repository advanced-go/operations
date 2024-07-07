package caseofficer1

import (
	"fmt"
	"github.com/advanced-go/operations/activity1"
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
	status := logActivity([]activity1.Entry{{AgentId: "agent-id"}})

	fmt.Printf("test: logActivity() -> [status:%v]\n", status)

	//Output:
	//test: logActivity() -> [status:OK]

}

func ExampleInsertAssignmentStatus() {
	msg := messaging.NewMessageWithStatus(messaging.ChannelStatus, "to", "from", "", core.StatusOK())
	status := insertAssignmentStatus(msg)

	fmt.Printf("test: insertAssignmentStatus() -> [status:%v]\n", status)

	//Output:
	//test: insertAssignmentStatus() -> [status:OK]

}

func ExampleUpdateAssignments() {
	entries, status := updateAssignments(core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "",
		Host:       "",
		InstanceId: "",
	})
	fmt.Printf("test: updateAssignments() -> [status:%v] [entries:%v]\n", status, entries)

	//Output:
	//test: updateAssignments() -> [status:OK] [entries:[{us-central1 c  www.host1.com 2024-06-10 09:00:35 +0000 UTC} {us-central1 c  www.host2.com 2024-06-10 09:00:35 +0000 UTC}]]

}

func ExampleNewAgent2() {
	origin := core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "",
		Host:       "www.host1.com",
		InstanceId: "",
	}
	a := newAgent(access.IngressTraffic, origin, nil)
	fmt.Printf("test: newAgent(\"%v\") -> [%v]\n", access.IngressTraffic, a)

	a = newAgent(access.EgressTraffic, origin, nil)
	fmt.Printf("test: newAgent(\"%v\") -> [%v]\n", access.EgressTraffic, a)

	//Output:
	//test: newAgent("ingress") -> [ingress-controller1:us-central1.c.www.host1.com]
	//test: newAgent("egress") -> [egress-controller1:us-central1.c.www.host1.com]

}

func ExampleProcessAssignments() {
	origin := core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "",
		Host:       "www.host1.com",
		InstanceId: "",
	}

	c := newCaseAgent(time.Second*5, access.IngressTraffic, origin, nil)
	fmt.Printf("test: newCaseAgent() -> [status:%v]\n", c != nil)

	status := processAssignments(c, logActivity, updateAssignments, newAgent)
	fmt.Printf("test: processAssignments() -> [status:%v] [controllers:%v]\n", status, c.controllers.Count())

	//Output:
	//test: newCaseAgent() -> [status:true]
	//test: processAssignments() -> [status:OK] [controllers:2]

}
