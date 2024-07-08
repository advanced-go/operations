package logistics1

import (
	"fmt"
)

func ExampleAgentUri() {
	u := AgentUri("west")
	fmt.Printf("test: AgentUri() -> [%v]\n", u)

	//Output:
	//test: AgentUri() -> [logistics1:west]

}

func ExampleNewAgent() {
	// a := NewAgent()
	fmt.Printf("test: newAgent() -> ")

	//Output:
	//test: newAgent() ->

}
