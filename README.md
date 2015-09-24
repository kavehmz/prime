prime
=========
[![Build Status](https://travis-ci.org/kavehmz/prime.svg)](https://travis-ci.org/kavehmz/prime)
[![Coverage Status](https://coveralls.io/repos/kavehmz/prime/badge.svg?branch=master&service=github)](https://coveralls.io/github/kavehmz/prime?branch=master)

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
