package landscape1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

func ExamplePut() {
	_, status := put[core.Output, Entry](nil, nil, partitionResource, "", nil, nil)
	fmt.Printf("test: put(nil,h,nil) -> [status:%v] [count:%v]\n", status, len(entryData))

	_, status = put[core.Output, Entry](nil, nil, partitionResource, "", []Entry{{Partition: 45}}, nil)
	fmt.Printf("test: put(nil,h,[]Entry) -> [status:%v] [count:%v]\n", status, len(entryData))

	//Output:
	//test: put(nil,h,nil) -> [status:Invalid Content [error: no entries found]] [count:6]
	//test: put(nil,h,[]Entry) -> [status:OK] [count:7]

}
