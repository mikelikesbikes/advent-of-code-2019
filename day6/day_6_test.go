package main

import (
	"strings"
	"testing"
)

type TH struct {
	pw string
	e  bool
}

func TestUOM(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	orbits := strings.Split(input, "\n")
	uom := buildUOM(orbits)
	if uom.m["B"] != "COM" {
		t.Fatal("failed to parse COM")
	}
}

func TestUOMChecksum(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	orbits := strings.Split(input, "\n")
	uom := buildUOM(orbits)
	if cs := uom.checksumAt("D"); cs != 3 {
		t.Fatalf("expected checksum at D to equal 3, got %v", cs)
	}
	if cs := uom.checksumAt("L"); cs != 7 {
		t.Fatalf("expected checksum at L to equal 7, got %v", cs)
	}
	if cs := uom.Checksum(); cs != 42 {
		t.Fatalf("expected checksum to equal 42, got %v", cs)
	}
}

func TestTransferPath(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
	orbits := strings.Split(input, "\n")
	uom := buildUOM(orbits)

	expected := []string{"K", "J", "E", "D", "I"}

	if tp := uom.TransferPath("YOU", "SAN"); !Equal(tp, expected) {
		t.Fatalf("expected transfer path to be %v, got %v", expected, tp)
	}
}

//func TestValidPassword(t *testing.T) {
//	for _, th := range []TH{
//		{pw: "111111", e: true},
//		{pw: "223450", e: false},
//		{pw: "123789", e: false},
//	} {
//
//		actual := ValidPassword(th.pw)
//		if actual != th.e {
//			t.Fatalf("expected %v, got %v", th.e, actual)
//		}
//	}
//}

func Equal(a, b []string) bool {
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
