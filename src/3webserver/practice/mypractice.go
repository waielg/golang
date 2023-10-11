package main

import (
	"fmt"
	"math"
)

func sumgbt(x1 int, x2 int) int {
	return x1 + x2
}

func factorial(n int) int {
	result := 1

	// Calculate factorial using an iterative loop
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func primecheck(n int) bool {
	if n <= 1 {
		return false // Numbers less than or equal to 1 are not prime.
	}
	// Check divisibility from 2 up to the square root of n.
	maxDivisor := int(math.Sqrt(float64(n)))
	for i := 2; i <= maxDivisor; i++ {
		if n%i == 0 {
			return false // Found a divisor, so n is not prime.
		}
	}

	return true // If no divisors found, n is prime.
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(sumgbt(646, 6476))
	fmt.Println(factorial(4))
	fmt.Println(primecheck(44))

}
