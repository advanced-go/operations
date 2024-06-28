package landscape1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleGet() {
	q := "status=active"
	entries, _, status := get[core.Output, Entry](nil, nil, uri.BuildValues(q), partitionResource, "", nil)
	fmt.Printf("test: get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, len(entries))

	q = "assigned-region=west&order=desc"
	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q), partitionResource, "", nil)
	fmt.Printf("test: get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, entries)

	//Output:
	//test: get("status=active") -> [status:OK] [entries:6]
	//test: get("assigned-region=west&order=desc") -> [status:OK] [entries:[{3 us-south1 c  ingress active  2024-06-10 09:00:35 +0000 UTC ingress-case-class1 west} {2 us-west1 b  egress active  2024-06-10 09:00:35 +0000 UTC egress-case-class1 west} {1 us-west1 a  ingress active  2024-06-10 09:00:35 +0000 UTC ingress-case-class1 west}]]

}
