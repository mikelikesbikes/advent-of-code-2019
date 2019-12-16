package main

import (
	//"fmt"
	"testing"
)

var input = []string{
	"12345678",
}

func TestFFT(t *testing.T) {
	input := "12345678"
	expected := "48226158"
	if actual := FFT(input, 1); actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}
	expected = "34040438"
	if actual := FFT(input, 2); actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}
	expected = "03415518"
	if actual := FFT(input, 3); actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}
	expected = "01029498"
	if actual := FFT(input, 4); actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}

	input = "80871224585914546619083218645595"
	expected = "24176176"
	if actual := FFT(input, 100)[0:8]; actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}
	input = "19617804207202209144916044189917"
	expected = "73745418"
	if actual := FFT(input, 100)[0:8]; actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}
	input = "69317163492948606335995924319873"
	expected = "52432133"
	if actual := FFT(input, 100)[0:8]; actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	input := "03036732577212944063491565474664"
	expected := "84462026"
	if actual := Decode(input, 100); actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}
	input = "02935109699940807407585447034323"
	expected = "78725270"
	if actual := Decode(input, 100); actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
	}
	input = "03081770884921959731165446850517"
	expected = "53553731"
	if actual := Decode(input, 100); actual != expected {
		t.Fatalf("expected fft to be %v, got %v", expected, actual)
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
