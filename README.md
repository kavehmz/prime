Prime
=========
[![Go Lang](http://kavehmz.github.io/static/gopher/gopher-front.svg)](https://golang.org/)
[![Build Status](https://travis-ci.org/kavehmz/prime.svg)](https://travis-ci.org/kavehmz/prime)
[![Coverage Status](https://coveralls.io/repos/kavehmz/prime/badge.svg?branch=master&service=github)](https://coveralls.io/github/kavehmz/prime?branch=master)
[![GoDoc](https://godoc.org/gopkg.in/kavehmz/prime.v1?status.svg)](https://godoc.org/gopkg.in/kavehmz/prime.v1)

This is a [Go](http://golang.org) library to produce prime numbers using all available cpu cores.


## Installation

```bash
$ go get github.com/kavehmz/prime
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/kavehmz/prime"
)

func main() {
	p := prime.Primes(1000000)
	fmt.Println("Number of primes:", len(p))
}
```
# Algorithm
To find more about different methods to find a range of prime numbers you can look at following pages:

* [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes) This is a more memory demanding method but faster by far for larger numbers.
* [Trial division](https://en.wikipedia.org/wiki/Trial_division) Easier to understand and less memory consuming.

# Performance
Performance depends on the size of max number. But as an example, it needs about 3ms to produce the first 1,000,000 prime numbers.


```bash
$ go test -bench .  
PASS
BenchmarkPrimes-4	     300	   3948274 ns/op
ok  	github.com/kavehmz/prime	1.618s
```

x          |time
-----------|------
1,000,000  |0.003s
10,000,000 |0.081s
100,000,000|1.292s

These calculations are done on a Mac 3.1 GHz Intel Core i7.

# Why
I used this simple library mainly to learn about Go language and Go standards. Also to use it for solving problems in https://projecteuler.net/

But I think it might be useful as a relatively fast implementation of prime numbers generator in Go.
