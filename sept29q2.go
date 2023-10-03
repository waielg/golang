package main

import (
	"fmt"
	"strconv"
)

func main() {
    a := []int{11, 22} // should be 2200
    result := concnew(a)
    fmt.Println(result)
}


func concnew(a []int) int64 {
var sumnew int64 = 0
n := len(a)
for i := 0; i < n; i++ {
	for j := 0; j < n; j++ {
		prod := fmt.Sprintf("%d%d", a[i], a[j])
		val := parseInt(prod)
		sumnew += int64(val)
	}
}

return sumnew
}

func parseInt(s string) int {
    val, err := strconv.Atoi(s)
    if err != nil {
        // Handle the error if conversion fails (not expected in this case)
        return 0
    }
    return val
}



	