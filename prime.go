// Package prime provides functionality to produce prime numbers using all
// available cpu cores. https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
// can be an starting point to find more information about how to calculate
// prime numbers.
//
// The method used in Primes function is Segmented sieve. Segmenting will
// Reduce memory requirement of process.
// The space complexity of the algorithm is O(√n).
package prime

import (
	"math"
	"runtime"
)

func fill(nums []bool, i uint64, max uint64) {
	a := 3 * i
	for a <= max {
		nums[a/2] = true
		a = a + 2*i
	}
}

func goFill(nums []bool, i uint64, max uint64, next chan bool) {
	fill(nums, i, max)
	<-next
}

// SieveOfEratosthenes returns a slice of all prime numbers equal or lower than max using Sieve Of Eratosthenes.
// This is without segmenting.
func SieveOfEratosthenes(n uint64) []uint64 {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	next := make(chan bool, cores)

	var nums = make([]bool, n/2+1)
	m := uint64(math.Sqrt(float64(n)))

	for i := uint64(3); i <= m; i = i + 2 {
		if nums[i/2] == false {
			go goFill(nums, i, n, next)
			next <- true
		}
	}

	for i := 0; i < cores; i++ {
		next <- true
	}

	var ps []uint64
	if n >= 2 {
		ps = append(ps, 2)
	}
	for i := uint64(3); i <= n; i = i + 2 {
		if nums[i/2] == false {
			ps = append(ps, i)
		}
	}
	return ps
}

func fillSegments(n uint64, basePrimes []uint64, allPrimes *[]uint64, segSize uint64, segNum uint64, next chan bool, nextTurn []chan bool) {
	cseg := make([]bool, segSize)
	for i := 0; i < len(basePrimes); i++ {
		jMax := segSize * (segNum + 1) / basePrimes[i]
		for j := (segSize * segNum) / basePrimes[i]; j < jMax; j++ {
			sn := (j + 1) * basePrimes[i]
			cseg[sn-segSize*segNum-1] = true
		}
	}

	// This waiting for turn is to avoid sorts at the end.
	// Sorts are much more expensive than this wait even for a
	// mostly sorted list.
	if segNum > 1 {
		<-nextTurn[segNum]
	}

	for i := uint64(0); i < segSize; i++ {
		if !cseg[i] && segSize*segNum+i+1 <= n {
			*allPrimes = append(*allPrimes, segSize*segNum+i+1)
		}
	}
	if int(segNum)+1 < len(nextTurn) {
		nextTurn[segNum+1] <- true
	}
	<-next

}

// Primes is using Segmented sieve. This method will reduce memory usae of Sieve of Eratosthenes considerably.
// besides memory allocation for Prime numbers slice, there is only O(sqrt(n)) extra memory required for the operation
// You can learn more about it in https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes.
func Primes(n uint64) (allPrimes []uint64) {
	if uint64(math.Log(float64(n))-1) == 0 {
		return SieveOfEratosthenes(n)
	}

	// There is a function pi(x) in math that will returns approximate number of prime numbers below n.
	allPrimes = make([]uint64, 0, n/uint64(math.Log(float64(n))-1))
	segSize := uint64(math.Sqrt(float64(n)))

	basePrimes := SieveOfEratosthenes(segSize)
	allPrimes = append(allPrimes, basePrimes...)

	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	next := make(chan bool, cores)
	var nextTurn []chan bool
	nextTurn = make([]chan bool, n/segSize+1)
	for i := uint64(0); i < n/segSize+1; i++ {
		nextTurn[i] = make(chan bool)
	}
	for segNum := uint64(1); segNum <= n/segSize; segNum++ {
		go fillSegments(n, basePrimes, &allPrimes, segSize, segNum, next, nextTurn)
		next <- true
	}
	for i := 0; i < cores; i++ {
		next <- true
	}

	return allPrimes
}
