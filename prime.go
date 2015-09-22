package prime

import (
	"math"
	"runtime"
	"sync"
)

var p []uint64
var mutex = &sync.Mutex{}

func is_prime(n uint64) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	m := uint64(math.Sqrt(float64(n)))
	for i := uint64(3); i <= m; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func chk(min uint64, max uint64, ready chan bool) {
	var t []uint64
	for i := min; i <= max; i++ {
		if is_prime(i) {
			t = append(t, i)
		}
	}
	mutex.Lock()
	p = append(p, t...)
	mutex.Unlock()
	<-ready

}

func Primes(max uint64) []uint64 {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ready := make(chan bool, runtime.NumCPU())

	for i := uint64(2); i <= max; i = i + 10000 {
		if i+9999 <= max {
			go chk(i, i+9999, ready)
		} else {
			go chk(i, max, ready)
		}
		ready <- true
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		ready <- true
	}

	return p
}
