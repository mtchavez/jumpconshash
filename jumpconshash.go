package jumpconshash

// OFFSET constant for hashing
var OFFSET uint64 = 2862933555777941757

// MAXINT is the max int size allowed
var MAXINT = float64(int64(1) << 31)

// Hash takes an uint64 key and a total number of buckets to consider to hash
// the key into one of the buckets consistently
func Hash(key uint64, buckets int) (b int64) {
	var j int64
	b = -1
	for j < int64(buckets) {
		b = j
		key = key*OFFSET + 1
		j = int64(float64(b+1) * (MAXINT / float64((key>>33)+1)))
	}
	return
}
