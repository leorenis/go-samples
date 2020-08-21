package popcount

import "testing"

// To run this benchmark change to dir ch1 and run:
// $  go test -bench=.

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountClearRightMost is
func PopCountClearRightMost(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1 // Clear the bit most rignt
		count++
	}
	return count
}

func BenchmarkClearRightMost(b *testing.B) {
	bench(b, PopCountClearRightMost)
}
