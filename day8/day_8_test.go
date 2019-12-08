package main

import "testing"

func TestHelloWorld(t *testing.T) {
	image := Image{height: 2, width: 3, data: "123456789012"}
	if image.Checksum() != 1 {
		t.Fatal("failed to parse image data")
	}
}

func TestFlatten(t *testing.T) {
	image := Image{height: 2, width: 2, data: "0222112222120000"}
	expected := []int{0, 1, 1, 0}
	if actual := image.flatten(); !Equal(actual, expected) {
		t.Fatalf("expected %v, got %v", expected, actual)
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
