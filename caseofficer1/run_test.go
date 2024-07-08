package caseofficer1

import (
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	fmt2 "github.com/advanced-go/stdlib/fmt"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

func testLog(_ context.Context, agentId string, content any) *core.Status {
	fmt.Printf("test: %v : %v : %v\n", fmt2.FmtRFC3339Millis(time.Now().UTC()), agentId, content)
	return core.StatusOK()
}

type testAgent struct{}

func newTestAgent() *testAgent {
	return new(testAgent)
}

func (t *testAgent) Uri() string { return "" }

func (t *testAgent) Message(m *messaging.Message) {
	fmt.Printf("%v", m)
}

func (t *testAgent) Run() {}

func (t *testAgent) Shutdown() {}

func ExampleRun() {

	fmt.Printf("test: run() -> [%v]\n", "")

	//Output:
	//fail

}
