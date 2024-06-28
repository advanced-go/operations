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

	q = "assigned-region=west&status=active&traffic=ingress"
	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q), partitionResource, "", nil)
	fmt.Printf("test: get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, entries)

	//Output:
	//test: get("status=active") -> [status:OK] [entries:6]
	//test: get("assigned-region=west&status=active&traffic=ingress") -> [status:OK] [entries:[{1 us-west1 a  ingress active  2024-06-10 09:00:35 +0000 UTC ingress-case-class1 west} {5 us-central1 c  ingress active  2024-06-10 09:00:35 +0000 UTC ingress-case-class1 west}]]

}
