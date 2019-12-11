package main

import (
	"fmt"
	"testing"
)

func TestRelativeBaseOffset(t *testing.T) {
	p := NewProgram([]int{109, 19, 204, -34})
	p.relativeBase = 2000

	p.Step()

	if p.relativeBase != 2019 {
		t.Fatalf("expected relativeBase to be 2019, got %v", p.relativeBase)
	}

	p.Step()

	out := <-p.out
	if out != 0 {
		t.Fatalf("expected 0")
	}
}
func TestQuine(t *testing.T) {
	input := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	p := NewProgram(input)

	go func() {
		defer close(p.out)
		p.Run()
	}()

	var output []int
	for i := range p.out {
		output = append(output, i)
	}
	if !Equal(input, output) {
		t.Fatalf("expected %v, got %v", input, output)
	}
}
func Test16Digit(t *testing.T) {
	input := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	p := NewProgram(input)

	go func() {
		defer close(p.out)
		p.Run()
	}()

	actual := <-p.out
	if len(fmt.Sprintf("%v", actual)) != 16 {
		t.Fatalf("expected 16 digit number, got %v", actual)
	}
}
func TestEcho16Digit(t *testing.T) {
	input := []int{104, 1125899906842624, 99}
	p := NewProgram(input)

	go func() {
		defer close(p.out)
		p.Run()
	}()

	actual := <-p.out
	if actual != input[1] {
		t.Fatalf("expected %v, got %v", input[1], actual)
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
