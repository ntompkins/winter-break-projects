package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// findPrimes will calculate all primes up to a certain point and return the list
func findPrimes(x int) []int {
    primes := []int{2}
    for i := 2; i < x; i++ {
        if isPrime(i) {
            primes = append(primes, i)
        }
    }
    return primes
}

// Determines if a number is prime or not
func isPrime(x int) bool {

	// Check if x is even
	if x % 2 == 0 {
		return false
	}

	// Run through all numbers
    for i := x-1; i > 1; i-- {
        if x % i == 0 {
            return false
        }
    }
    return true
}

// Linear search to find max value in list
func max(list []int) int {
    m := list[0]
    for i := 0; i < len(list); i++ {
        if list[i] > m {
            m = list[i]
        }
    }
    return m
}

// Find prime sums for numbers
func goldbachSums(x int, primes []int) []int {

	var start int
	var sums []int

	// Start by finding closest prime to our number
	for i := 0; i < len(primes); i++ {
		if primes[i] < x {
			start = i
		}
	}
	// Begin at start and work through list of primes
	for small := 0; small < start+1; small++ {
		for big := start; big >= small; big-- {
			if primes[big] + primes[small] == x {
				sums = append(sums, primes[small])
				sums = append(sums, primes[big])
			}
		}
	}
	return sums
}

func main() {

    // Create scanner for input
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())

    // Create list to hold numbers
    input := make([]int, n)

    // Insert numbers into list
    for i := 0; i < n; i++{
        scanner.Scan()
        input[i], _ = strconv.Atoi(scanner.Text())
    }

    // Find primes up to our largest input number
    primes := findPrimes(max(input))

	var sums []int
	//var builder strings.Builder

    for i := 0; i < len(input); i++ {
		sums = goldbachSums(input[i], primes)
		fmt.Printf("%d has %d representation(s)\r\n", input[i], len(sums)/2)

		for j := 0; j < len(sums); j+=2 {
			fmt.Printf("%d+%d\r\n", sums[j], sums[j+1])
		}

		fmt.Println()
    }
}

