package activity1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

func ExampleLog() {
	status := Log(nil, "agent-id", core.Origin{Region: "us-central1", Zone: "c"})

	fmt.Printf("test: Log() -> [status:%v] [entry:%v]\n", status, lastEntry().Content)

	//Output:
	//test: Log() -> [status:OK] [entry:{us-central1 c   }]

}
