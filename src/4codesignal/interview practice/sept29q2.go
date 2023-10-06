package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []int{11, 22} // should be 6666
	result := concnew(a)
	fmt.Println(result)
}

func concnew(a []int) int64 {
	var sumnew int64 = 0
	n := len(a)
	for i := 0; i < n; i++ { // nested loop to iterate through a
		for j := 0; j < n; j++ {
			prod := fmt.Sprintf("%d%d", a[i], a[j]) // from int we created string "prod"
			val := parseInt(prod)                   // convert concatenated string back to int
			sumnew += int64(val)
		}
	}
	return sumnew
}

func parseInt(s string) int { // convert string to int
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
