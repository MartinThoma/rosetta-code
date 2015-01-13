package main

import "fmt"
import "sort"
import "math/rand"
import "time"

func main() {
    // Seed
    rand.Seed( time.Now().UTC().UnixNano())

    var slice []int
    var n = 10000000

    fmt.Printf("Start creating %v random integers...\n", n)
    for i := 0; i < n; i++ {
        slice = append(slice, rand.Intn(1000000000))
    }

    fmt.Printf("Start sorting %v random integers...\n", n)
    start := time.Now()
    sort.Sort(sort.IntSlice(slice))
    elapsed := time.Since(start)
    fmt.Printf("%s\n", elapsed)
}
