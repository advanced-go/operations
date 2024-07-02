package assignment1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleGet_Entry() {
	q := "region=*"
	entries, _, status := get[core.Output, Entry](nil, nil, uri.BuildValues(q), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, len(entries))

	q = "region=*&order=desc"
	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, entries)

	//Output:
	//test: Get("region=*") -> [status:OK] [entries:5]
	//test: Get("region=*&order=desc") -> [status:OK] [entries:[{us-west1 a  www.host1.com 2024-06-10 09:00:35 +0000 UTC} {us-west1 a  www.host2.com 2024-06-10 09:00:35 +0000 UTC} {us-central1 c  www.host1.com 2024-06-10 09:00:35 +0000 UTC} {us-central1 c  www.host2.com 2024-06-10 09:00:35 +0000 UTC} {us-central1 c  www.host3.com 2024-06-10 09:00:35 +0000 UTC}]]

}
