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
	t := time.Now().UnixNano()

	fmt.Println("number of primes (non-segmented-method):", len(prime.SieveOfEratosthenes(1000000000)))
	fmt.Println("number of primes (segmented-method):", len(prime.Primes(1000000)))
	fmt.Println("seconds it took:", float64(time.Now().UnixNano()-t)/1000000000)

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
