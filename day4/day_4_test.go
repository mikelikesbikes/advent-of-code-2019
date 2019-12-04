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

func TestValidPassword2(t *testing.T) {
	for _, th := range []TH{
		{pw: "112233", e: true},
		{pw: "123444", e: false},
		{pw: "111122", e: true},
	} {

		actual := ValidPassword2(th.pw)
		if actual != th.e {
			t.Fatalf("expected %v, got %v for %v", th.e, actual, th.pw)
		}
	}
}
