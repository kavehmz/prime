package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/kavehmz/prime"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var primeRange = flag.Uint64("primeRange", 1000000000, "Set the max prime number range")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	fmt.Println("number of primes (non-segmented-method):", len(prime.SieveOfEratosthenes(*primeRange)))
	fmt.Println("seconds it took:", time.Since(start))

	start = time.Now()
	fmt.Println("number of primes (segmented-method):", len(prime.Primes(*primeRange)))
	fmt.Println("seconds it took:", time.Since(start))

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}

}
