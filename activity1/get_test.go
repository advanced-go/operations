package activity1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleGet() {
	q := "region=*"
	entries, _, status := get[core.Output, Entry](nil, nil, uri.BuildValues(q), activityResource, "", nil)
	fmt.Printf("test: get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, len(entries))

	q = "region=*&order=desc"
	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q), activityResource, "", nil)
	fmt.Printf("test: get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, entries)

	//Output:
	//test: get("region=*") -> [status:OK] [entries:1]
	//test: get("region=*&order=desc") -> [status:OK] [entries:[{1 agent-id 2024-06-10 09:00:35 +0000 UTC testing 1-2-3}]]

}
