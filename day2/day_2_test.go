package main

import (
	"fmt"
	"testing"
)

func TestCompute(t *testing.T) {
	for _, e := range [][][]int{
		[][]int{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		[][]int{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		},
		[][]int{
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		[][]int{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	} {
		res := Compute(e[0])
		if !Equal(res, e[1]) {
			t.Fatal(fmt.Sprintf("expected %v, got %v", e[1], res))
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
