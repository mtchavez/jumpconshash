package jumpconshash

import (
	"math"
	"testing"
)

func TestHash_TestJumps(t *testing.T) {
	// Map of key to hash with an slice
	// of expected bucket to be returned for the
	// total buckets of (index + 1) in the slice
	jumps := map[uint64][]int64{
		0:              []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		1:              []int64{0, 0, 0, 0, 0, 0, 6, 6, 6, 6, 6, 6, 6, 6},
		3735928559:     []int64{0, 1, 2, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		math.MaxUint32: []int64{0, 0, 2, 2, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5},
	}
	for key, expectedJumps := range jumps {
		for bucket, expected := range expectedJumps {
			if hash := Hash(key, bucket+1); hash != expected {
				t.Errorf("[Fail] Hash of %+v for %+v buckets should be %+v but got %+v", key, bucket+1, expected, hash)
			}
		}
	}
}
