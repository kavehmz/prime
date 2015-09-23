package prime

import (
	"math"
	"runtime"
)

var ps []uint64

func fill(i uint64, max uint64, ready chan bool) {
	a := 2 * i
	for a <= max {
		ps[a] = 1
		a = a + i
	}
	<-ready
}

func Primes(max uint64) []uint64 {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ready := make(chan bool, runtime.NumCPU())

	ps = make([]uint64, max+1)
	m := uint64(math.Sqrt(float64(max)))
	for i := uint64(2); i <= m; i = i + 2 {

		if ps[i] == 0 {
			go fill(i, max, ready)
			ready <- true
		}
		if i == 2 {
			i = 1
		}

	}

	for i := 0; i < runtime.NumCPU(); i++ {
		ready <- true
	}

	var t []uint64
	m = uint64(len(ps))
	for i := uint64(2); i < uint64(len(ps)); i++ {
		if ps[i] == 0 {
			t = append(t, ps[i])
		}
	}
	return t
}
