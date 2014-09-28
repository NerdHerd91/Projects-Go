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
    e := computeE()

    // Format the value of e to the number of digits requested.
    fmt.Printf("The value of e to %v digits: %v\n", n, strconv.FormatFloat(e, 'f', n, 64)) 
}

func computeE() (float64) {
    return math.Exp(1)
}
