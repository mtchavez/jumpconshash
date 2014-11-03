# Jump Consistent Hash
[![Build Status](https://drone.io/github.com/mtchavez/jumpconshash/status.png)](https://drone.io/github.com/mtchavez/jumpconshash/latest)

Jump consistent hash implementation in Go.

## Background

Based on [this whitepaper](http://www.smallake.kr/wp-content/uploads/2014/08/1406.2294.pdf) from Google. Jump consistent hash takes an integer key and a total number of buckets to consider and hashes the key into one of the buckets. The algorithm achieves minimal memory consumption and fast hashing that has a small change in hashing as total buckets change.

## Docs

Docs are on [GoDoc](http://godoc.org/github.com/mtchavez/jumpconshash)

## Tests

Tests can be ran with `go test --cover`

## Usage

[Example usage](http://godoc.org/github.com/mtchavez/jumpconshash#example-Hash)

```go
// Hash takes the integer key you want
// to find the bucket it is hashed to
// and the total number of buckets to consider
totalBuckets := 100
bucket := Hash(123456789, totalBuckets)

// The bucket for the key is returned
fmt.Printf("Hash(123456789, 100) is bucket %+v", bucket)
```
