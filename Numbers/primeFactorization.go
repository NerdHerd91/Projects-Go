package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
)

func printPrimeFactorization(n int) {
    factors := []int{} // Create an empty list for our factors.
    for i := 2; i < n / i; i++ {
        for n % i == 0 {
            factors = append(factors, i)
            n /= i
        }
    }
    if n > 1 {
        factors = append(factors, n)
    }
    fmt.Println(factors)
}

func main() {
    // Check the number of arguments passed.
    if len(os.Args) < 2 {
        log.Fatal("An integer N must be provided as the first argument.")
    }

    // Attempt to parse the argument for n.
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    printPrimeFactorization(n)
}
