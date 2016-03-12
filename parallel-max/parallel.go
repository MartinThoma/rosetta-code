package main

// A parallel max version
// TODO: Doesn't work right now!
//
// This script should create a big array (slice?) of pseudo-random integers.
// The task is to find the maximum of it. To do so, the array should be split
// in several sub-arrays of similar size (without copying the content). Finding
// the maximum in the sub-arrays should be done in parallel. After that, the
// maximum of the complete array should be returned. Measure the time of that
// and the simple way to do it.

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func SimpleMax(numbers []int64) int64 {
	max_ := numbers[0]
	for _, element := range numbers {
		if element > max_ {
			max_ = element
		}
	}
	return max_
}

func getMax(numbers []int64, c chan int64) {
	max_ := numbers[0]
	for _, element := range numbers {
		if element > max_ {
			max_ = element
		}
	}
	c <- max_
}

func ParallelMax(numbers []int64) int64 {
	max_ := numbers[0]
	i := 0
	var c chan int64 = make(chan int64)
	for i != len(numbers) {
		go getMax(numbers[8*i:8*(i+1)], c)
		i += 8
	}

	// barrier?

	for len(c) > 0 {
		element := <-c
		if element > max_ {
			max_ = element
		}
	}

	return max_
}

func main() {
	fmt.Println("Hello World")
	rand.Seed(42)
	count := 100000000 // 200000000
	numbers := make([]int64, count, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Int63()
	}

	start := time.Now()
	log.Printf("Max: %d", SimpleMax(numbers))
	elapsed := time.Since(start)
	log.Printf("SimpleMax took %s", elapsed)

	start = time.Now()
	fmt.Println("Max: %i", ParallelMax(numbers))
	elapsed = time.Since(start)
	log.Printf("ParallelMax took %s", elapsed)
}
