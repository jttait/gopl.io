package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Exercise 2.3
func PopCount2(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

// Exercise 2.4
func PopCount3(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x & 1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

// Exercise 2.5
func PopCount4(x uint64) int {
	count := 0
	for x > 0 {
		x = x & (x-1)
		count++
	}
	return count
}
