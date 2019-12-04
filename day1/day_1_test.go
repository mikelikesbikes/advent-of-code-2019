package main

import (
	"fmt"
	"testing"
)

func TestFuelCounterUpper(t *testing.T) {
	for _, e := range [][]int{
		[]int{12, 2},
		[]int{14, 2},
		[]int{1969, 654},
		[]int{100756, 33583},
	} {
		v := FuelCounterUpper(e[0])
		if v != e[1] {
			t.Fatal(fmt.Sprintf("expected %v, got %v", e[1], v))
		}
	}
}

func TestFuelDoubleChecker(t *testing.T) {
	for _, e := range [][]int{
		[]int{12, 2},
		[]int{14, 2},
		[]int{1969, 966},
		[]int{100756, 50346},
	} {
		v := FuelDoubleChecker(e[0])
		if v != e[1] {
			t.Fatal(fmt.Sprintf("expected %v, got %v", e[1], v))
		}
	}
}
