package prime

import (
	"math"
	"runtime"
)

func fill(nums []bool, i uint64, max uint64) {
	a := 2 * i
	for a <= max {
		nums[a] = true
		a = a + i
	}
}

func go_fill(nums []bool, i uint64, max uint64, ready chan bool) {
	fill(nums, i, max)
	<-ready
}

func Primes(max uint64) []uint64 {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ready := make(chan bool, runtime.NumCPU())

	var nums = make([]bool, max+1)
	m := uint64(math.Sqrt(float64(max)))
	for i := uint64(2); i <= m; i = i + 2 {

		if nums[i] == false {
			go go_fill(nums, i, max, ready)
			ready <- true
		}
		if i == 2 {
			i = 1
		}

	}

	for i := 0; i < runtime.NumCPU(); i++ {
		ready <- true
	}

	var ps []uint64
	m = uint64(len(nums))
	for i := uint64(2); i < uint64(len(nums)); i++ {
		if nums[i] == false {
			ps = append(ps, i)
		}
	}
	return ps
}
