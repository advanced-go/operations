package caseofficer1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
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
