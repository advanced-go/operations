package common

import (
	"fmt"
	"github.com/advanced-go/stdlib/uri"
)

const (
	q1 = "region=us-west-1&zone=usw1-az1&host=www.host1.com"
	q2 = "region=us-west-1&zone=usw1-az2&host=www.host2.com"
	q3 = "region=us-west-2&zone=usw2-az3&host=www.host1.com"
	q4 = "region=us-west-2&zone=usw2-az4&host=www.host2.com"
)

func ExampleOrder_Entry() {
	q := ""
	result := Order(nil, entryData)
	fmt.Printf("test: Order(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), result)

	q = "order=desc"
	result = Order(uri.BuildValues(q), entryData)
	fmt.Printf("test: Order(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), result)

	//Output:
	//test: Order("") -> [cnt:4] [result:[{1 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az1  www.host1.com} {2 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az2  www.host2.com} {3 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az3  www.host1.com} {4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com}]]
	//test: Order("order=desc") -> [cnt:4] [result:[{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com} {3 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az3  www.host1.com} {2 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az2  www.host2.com} {1 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az1  www.host1.com}]]

}

func ExampleTop_Entry() {
	q := ""
	result := Top(nil, entryData)
	fmt.Printf("test: Top(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), len(result))

	q = "top=2"
	result = Top(uri.BuildValues(q), entryData)
	fmt.Printf("test: Top(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), len(result))

	//Output:
	//test: Top("") -> [cnt:4] [result:4]
	//test: Top("top=2") -> [cnt:4] [result:2]

}

func ExampleFilterT_Entry() {
	entries, status := FilterT[Entry](uri.BuildValues(q1), entryData, ValidEntry)
	fmt.Printf("test: FilterT[Entry](\"%v\") -> [status:%v] [entries:%v]\n", q1, status, entries)

	//Output:
	//test: FilterT[Entry]("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [status:OK] [entries:[{1 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az1  www.host1.com}]]

}
