package popcount

import (
	"testing"
)

// To run this benchmark change to dir ch1 and run:
// $  go test -bench=.

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	/*j := 20
	println(j & 20)*/
}

func PopCountShiftValue(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 1 {
			count++
		}
		x >>= 1
	}
	return count
}

// BenchmarkShiftValue is
func BenchmarkShiftValue(b *testing.B) {
	bench(b, PopCountShiftValue)
}
