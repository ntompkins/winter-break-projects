package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(x int, y int, n int) string {

	// Validate input
	if x < 1 || x > y || n > 100 {
		return "Error: your input is invalid."
	}

	// Build results string
	result := ""
	for i := 1; i < n+1; i++ {

		// Determine what to add to string
		if i%x == 0 {
			result += "Fizz"
		}
		if i%y == 0 {
			result += "Buzz"
		}
		if i%y != 0 && i%x != 0 {
			result += strconv.Itoa(i)
		}

		// Add newline char
		result += "\n"
	}

	return result
}

func main() {

	// Read input from user
	var x, y, n int
	fmt.Scanf("%d %d %d", &x, &y, &n)

	// Call fizzbuzz and print result
	fmt.Print(fizzbuzz(x, y, n))
}
