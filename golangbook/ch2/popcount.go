package ch2

// pc[i] is the population of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	j := 1
	println(j & 1)
}
