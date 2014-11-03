package jumpconshash

var OFFSET uint64 = 2862933555777941757
var MAXINT float64 = float64(int64(1) << 31)

func Hash(key uint64, buckets int) (b int64) {
	var j int64 = 0
	b = -1
	for j < int64(buckets) {
		b = j
		key = key*OFFSET + 1
		j = int64(float64(b+1) * (MAXINT / float64((key>>33)+1)))
	}
	return
}
