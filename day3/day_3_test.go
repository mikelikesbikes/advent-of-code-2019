package main

import (
	"testing"
)

type TH struct {
	w1, w2 string
	d      int
}

func TestDistanceToCross(t *testing.T) {
	for _, e := range []TH{
		{
			w1: "R8,U5,L5,D3",
			w2: "U7,R6,D4,L4",
			d:  6,
		},
		{
			w1: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			w2: "U62,R66,U55,R34,D71,R55,D58,R83",
			d:  159,
		},
		{
			w1: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			w2: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			d:  135,
		},
	} {
		ad := DistanceToCross(e.w1, e.w2)
		if ad != e.d {
			t.Fatalf("expected %v, got %v", e.d, ad)
		}
	}
}

func TestEarliestCross(t *testing.T) {
	for _, e := range []TH{
		{
			w1: "R8,U5,L5,D3",
			w2: "U7,R6,D4,L4",
			d:  30,
		},
		{
			w1: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			w2: "U62,R66,U55,R34,D71,R55,D58,R83",
			d:  610,
		},
		{
			w1: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			w2: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			d:  410,
		},
	} {
		ad := EarliestCross(e.w1, e.w2)
		if ad != e.d {
			t.Fatalf("expected %v, got %v", e.d, ad)
		}
	}
}
