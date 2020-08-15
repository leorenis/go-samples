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

// PopcountTableLoop is Ex.2.3
func PopcountTableLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i))])
	}
	return sum
}

// BenchmarkTableLoop is
func BenchmarkTableLoop(b *testing.B) {
	bench(b, PopcountTableLoop)
}
