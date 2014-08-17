package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
)

// This function prints the first N fibonacci values.
// It takes in two seed values and the numeber of terms to compute.
func printNFib(f1, f2, n int) {
    // Start by printing the seed values (comma seperated).
    fmt.Printf("[%v, %v", f1, f2)

    // Print the next n-2 values of the fibonacci sequence.
    for i:= 2; i < n; i++ {
        sum := f1 + f2
        fmt.Printf(", %v", sum)
        f1 = f2
        f2 = sum
    }

    // Close the bracket and print a new line for the series.
    fmt.Printf("]\n")
}

func main() {
    // Formula for computing the fibonacci sequence is:
    // F_n = F_n-1 + F_n-2

    // Use seed values of [0,1].
    var f1, f2 int = 0, 1

    if len(os.Args) < 2 {
        log.Fatal("An integer N must be provided as the first argument.")
    }

    // Attempt to parse the argument for n.
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    printNFib(f1, f2, n)
}
