package main

func main() {

}

func solution(a []int) bool {

	// initialize and declare
	n := len(a)
	b := make([]int, n)

	// notice small arrays are true
	if n < 2 {
		return true
	}

	/*
		b[0] = a[0]
		b[1] = a[len(a)-1]
		b[2] = a[1]
		b[3] = a[len(a)-2]
		b[4] = a[2]
		b[5] = a[len(a)-3]
		b[6] = a[3]
		KEY: this is really just b[2*i] = a[i] and b[2*i+1] = a[n-1-i]
	*/

	// split into even and odd indexes
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			b[i] = a[i/2]
		} else {
			b[i] = a[n-1-i/2]
		}
	}

	for i := 1; i < len(b); i++ {
		if b[i] < b[i-1] {
			return false
		}
	}
	return true
}
