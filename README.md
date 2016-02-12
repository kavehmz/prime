Prime
=========
[![Go Lang](http://kavehmz.github.io/static/gopher/gopher-front.svg)](https://golang.org/)
[![GoDoc](https://godoc.org/github.com/kavehmz/prime?status.svg)](https://godoc.org/github.com/kavehmz/prime)
[![Build Status](https://travis-ci.org/kavehmz/prime.svg?branch=master)](https://travis-ci.org/kavehmz/prime)
[![Coverage Status](https://coveralls.io/repos/kavehmz/prime/badge.svg?branch=master&service=github)](https://coveralls.io/github/kavehmz/prime?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kavehmz/prime)](https://goreportcard.com/report/github.com/kavehmz/prime)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kavehmz/prime)

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
BenchmarkPrimes-4	     500	   3181972 ns/op
ok  	github.com/kavehmz/prime	1.618s
```

x            |time
-------------|------
1,000,000    |0.003s
10,000,000   |0.035s
100,000,000  |0.642s
1,000,000,000|8.253s

These calculations are done on a 3.1GHz Dual-core Intel Core i7.

# Why
I used this simple library mainly to learn Go language, Go standards and for solving problems in https://projecteuler.net/

It can also be useful as a relatively fast implementation of prime numbers generator in Go.
