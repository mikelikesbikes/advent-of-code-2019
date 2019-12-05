package main

import (
	"testing"
)

type TH struct {
	pw string
	e  bool
}

func TestValidPassword(t *testing.T) {
	for _, th := range []TH{
		{pw: "111111", e: true},
		{pw: "223450", e: false},
		{pw: "123789", e: false},
	} {

		actual := ValidPassword(th.pw)
		if actual != th.e {
			t.Fatalf("expected %v, got %v", th.e, actual)
		}
	}
}
