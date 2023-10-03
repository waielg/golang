package main

import "fmt"

// A list of different functinns

func adder(param1 int, param2 int) int {
	return param1 + param2
}

func decade(year int) int {
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

func largestproduct(inputArray []int) int {
	//define max
	max := inputArray[0] * inputArray[1]

	for i := 1; i < len(inputArray)-1; i++ {
		if max < inputArray[i]*inputArray[i+1] {
			max = inputArray[i] * inputArray[i+1]
		}
	}

	return max
}

func additionalstatues(statues []int) int {
	min := statues[0]
	max := statues[0]

	//identify range
	for val := range statues {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return (max - min + 1) - len(statues)
}

func removeoneintstrictlyincreasing(sequence []int) bool {
	removed := false
	j := 0
	for i := 1; i < len(sequence); i++ {
		if sequence[j] < sequence[i] {
			j = i
		} else {
			if removed {
				return false
			}
			if j == 0 {
				j = i
			}
			if sequence[i] > sequence[j-1] {
				j = i
			}
			removed = true
		}
	}
	return true
}

func matrixghostgame(matrix [][]int) int {

	sum := 0
	for x := 0; x < len(matrix[0]); x++ {
		for y := 0; y < len(matrix) && matrix[y][x] != 0; y++ {
			sum += matrix[y][x]
		}
	}

	return sum
}

func main() {
	fmt.Println("Hello, World im Waiel", palindrome("car"))
}