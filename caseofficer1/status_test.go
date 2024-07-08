package caseofficer1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
	"time"
)

func ExampleRunStatus() {
	origin := core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "",
		Host:       "www.host1.com",
		InstanceId: "",
	}
	msg := messaging.NewControlMessage("to", "from", messaging.ShutdownEvent)
	fmt.Printf("test: NewMessage() -> %v\n", msg.Event())

	c := newAgent(time.Second*1, access.IngressTraffic, origin, newTestAgent())
	go runStatus(c, testLog, insertAssignmentStatus)

	status := core.NewStatusError(http.StatusTeapot, errors.New("teapot error"))
	c.statusC <- messaging.NewMessageWithStatus(messaging.ChannelStatus, "to", "from", "event:status", status)
	time.Sleep(time.Second * 1)

	c.statusCtrlC <- msg
	time.Sleep(time.Second * 3)

	//fmt.Printf("test: InsertStatus() -> [entry:%v]")

	//Output:
	//fail
}
