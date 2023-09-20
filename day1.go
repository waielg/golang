// *******DAY11******
package main

import "fmt"

func main2() {
	fmt.Println("hello world")
}

// add
func solution1(param1 int, param2 int) int {
	return param1 + param2
}

// century
func solution2(year int) int {
	return (year + 99) / 100
}

// reverse boolean
func solution3(inputString string) bool {
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-1-i] {
			return false
		}
	}
	return true
}
