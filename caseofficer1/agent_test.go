package caseofficer1

import "fmt"

func ExampleAgentUri() {
	u := AgentUri("ingress", "us-central1", "c", "sub-zone")
	fmt.Printf("test: AgentUri() -> [%v]\n", u)

	u = AgentUri("egress", "us-west1", "a", "")
	fmt.Printf("test: AgentUri() -> [%v]\n", u)

	//Output:
	//test: AgentUri() -> [case-officer:ingress.us-central1.c.sub-zone]
	//test: AgentUri() -> [case-officer:egress.us-west1.a]

}

func ExampleNewAgent() {
	// a := NewAgent()
	fmt.Printf("test: newAgent() -> ")

	//Output:
	//test: newAgent() ->

}
