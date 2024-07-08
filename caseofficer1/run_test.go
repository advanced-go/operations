package caseofficer1

import (
	"context"
	"fmt"
	"github.com/advanced-go/operations/assignment1"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	fmt2 "github.com/advanced-go/stdlib/fmt"
	"time"
)

func testLog(_ context.Context, agentId string, content any) *core.Status {
	fmt.Printf("test: activity1.Log() -> %v : %v : %v\n", fmt2.FmtRFC3339Millis(time.Now().UTC()), agentId, content)
	return core.StatusOK()
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

	status := processAssignments(c, assignment1.Update, newControllerAgent)
	fmt.Printf("test: processAssignments() -> [status:%v] [controllers:%v]\n", status, c.controllers.Count())

	//Output:
	//test: newAgent() -> [status:true]
	//test: processAssignments() -> [status:OK] [controllers:2]

}

func ExampleRun() {
	fmt.Printf("test: run() -> [%v]\n", "")

	//Output:
	//fail

}
