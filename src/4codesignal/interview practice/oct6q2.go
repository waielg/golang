package main

func main() {

}

func solution(arr []int) int64 {
	var result int64 = 0

	for i := 1; i < len(arr)-1; i++ {
		if (arr[i] > arr[i-1] && arr[i] > arr[i+1]) || (arr[i] < arr[i-1] && arr[i] < arr[i+1]) {
			result++
		}
	}

	return result + 1
}
