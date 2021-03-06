package popcount

import (
	"fmt"
	"testing"
)

// To run this benchmark change to dir ch1 and run:
// $  go test -bench=.

// pc[i] is the population of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	/*j := 20
	println(j & 20)*/
}

// ShowPopCount is
func ShowPopCount() {
	fmt.Printf("%d", popcount(1000))
}

func popcount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// BenchmarkTable is
func BenchmarkTable(b *testing.B) {
	bench(b, popcount)
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}
