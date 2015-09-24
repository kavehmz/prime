package prime

import (
	"testing"
)

func TestPrimes(t *testing.T) {
	p := Primes(1000000)
	if len(p) != 78498 {
		t.Error("Wrong number of prime numbers lower than 1M ", len(p))
	}
	if p[0] != uint64(2) {
		t.Error("1st prime number is not 2 ", p[0])
	}
	if p[1] != uint64(3) {
		t.Error("2nd prime number is not 2 ", p[1])
	}
	if p[25] != uint64(101) {
		t.Error("25th prime number is not 2 ", p[25])
	}
	if p[23423] != uint64(267391) {
		t.Error("23424th prime number is not 267391 ", p[23423])
	}
	p = Primes(100)
	if p[24] != uint64(97) {
		t.Error("23424th prime number is not 267391 ", p[24])
	}
	if len(p) != 25 {
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