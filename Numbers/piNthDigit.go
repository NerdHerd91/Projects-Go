package main

import (
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("An integer N must be provided as the first argument.")
    }

    // Attempt to parse the argument for n.
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    // Check the precision requested is within our limits.
    if n < 1 || n > 51 {
        log.Fatal("An integer N must be within the range of [1, 51] inclusive.")
    }
    pi := computePi()

    // Format the value of pi to the number of digits requested.
    fmt.Printf("The value of Pi to %v digits: %v\n", n, strconv.FormatFloat(pi, 'f', n, 64)) 
}

func computePi() (float64) {
    // Use Machin's Formula to compute Pi.
    x := float64(1) / 5
    y := float64(1) / 239
    return 16 * math.Atan(x) - 4 * math.Atan(y) 
}
