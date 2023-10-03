package main

import "fmt"

// A list of different functinns

func adder(param1 int, param2 int) int {
	return param1 + param2
}

func topofheadadder(p1 int, p2 int) int {
	return p1 + p2
}

func decade(year int) int {
	return (year + 99) / 100
}

func topofheaddecade(year int) int {
	return (year + 99) / 100
}


func palindrome(inputString string) bool {
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-1-i] {
			return false
		}
	}
	return true
}

func topofheadpalindrome(word string) bool {
	for i:=0; i<len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println("lets test out this function. answer is ", topofheadpalindrome("racecar"))
}