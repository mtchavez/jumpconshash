package jumpconshash

import (
	"fmt"
)

func ExampleHash() {
	// Hash takes the integer key you want
	// to find the bucket it is hashed to
	// and the total number of buckets to consider
	totalBuckets := 100
	bucket := Hash(123456789, totalBuckets)

	// The bucket for the key is returned
	fmt.Printf("Hash(123456789, 100) is bucket %+v", bucket)
	// Output:
	// Hash(123456789, 100) is bucket 34
}
