package main

import (
	"fmt"
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


func main() {
	fmt.Println("Hello World")
	fmt.Println(sumgbt(646,6476))
	fmt.Println(factorial(4))

}