package assignment1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func ExampleInsertStatus() {
	origin := core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "",
		Host:       "www.host1.com",
		InstanceId: "",
	}
	status := core.NewStatus(http.StatusTeapot)
	result := InsertStatus(nil, "agent-id", origin, status)
	fmt.Printf("test: InsertStatus() -> [status:%v] [content:%v]\n", result, lastStatus().Content)

	status = core.NewStatusError(http.StatusGatewayTimeout, errors.New("this is an example of error content"))
	result = InsertStatus(nil, "agent-id", origin, status)
	fmt.Printf("test: InsertStatus() -> [status:%v] [content:%v]\n", result, lastStatus().Content)

	status = core.NewStatus(http.StatusOK)
	status.Content = origin
	result = InsertStatus(nil, "agent-id", origin, status)
	fmt.Printf("test: InsertStatus() -> [status:%v] [content:%v]\n", result, lastStatus().Content)

	//Output:
	//test: InsertStatus() -> [status:OK] [content:I'm A Teapot]
	//test: InsertStatus() -> [status:OK] [content:this is an example of error content]
	//test: InsertStatus() -> [status:OK] [content:{us-central1 c  www.host1.com }]

}

func ExampleUpdate() {
	entries, status := Update(nil, "agent-id", core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "",
		Host:       "",
		InstanceId: "",
	})
	fmt.Printf("test: Update() -> [status:%v] [entries:%v]\n", status, entries)

	//Output:
	//test: Update() -> [status:OK] [entries:[{us-central1 c  www.host1.com test-agent 2024-06-10 09:00:35 +0000 UTC} {us-central1 c  www.host2.com test-agent 2024-06-10 09:00:35 +0000 UTC}]]

}
