package main

import (
	"testing"
)

func testconcnew(t *testing.T) {	// Test case: Input [11, 22],  result should be 6666
    a1 := []int{11, 22}
    expected1 := int64(6666)
    result1 := concnew(a1)
    	if result1 != expected1 {
        	t.Errorf("Test case failed: Expected %d, got %d", expected1, result1)
    }
}















