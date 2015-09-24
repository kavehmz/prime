package prime

import (
	"testing"
)

func TestPrimes(t *testing.T) {
	p := Primes(1000000)
	if len(p) != 78498 {
		t.Error("Wrong number of prime numbers lower than 1M ", len(p))
	}
	p = Primes(10)
	if len(p) != 4 {
		t.Error("Wrong number of prime numbers lower than 10 ", len(p))
	}
	p = Primes(1)
	if len(p) != 0 {
		t.Error("Edge case of 1 not correct ", len(p))
	}
	p = Primes(2)
	if len(p) != 1 {
		t.Error("Edge case of 2 not correct ", len(p))
	}
}

func TestFill(t *testing.T) {
	p := make([]uint64, 101)
	fill(p, 2, 100)
	var c uint64
	for _, v := range p {
		if v == 1 {
			c++
		}
	}
	if c != 49 {
		t.Error("Filled cells are wrong ", c)
	}
}
