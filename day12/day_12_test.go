package main

import (
	//"fmt"
	"testing"
)

func TestStep(t *testing.T) {
	moons := []*Moon{
		NewMoon(-1, 0, 2),
		NewMoon(2, -10, -7),
		NewMoon(4, -8, 8),
		NewMoon(3, 5, -1),
	}

	initialVel := V{0, 0, 0}
	if moons[0].vel != initialVel ||
		moons[1].vel != initialVel ||
		moons[2].vel != initialVel ||
		moons[3].vel != initialVel {
		t.Fatal("didn't build moons with 0 velocity")
	}

	Step(moons)
	if moons[0].pos != (V{2, -1, 1}) ||
		moons[1].pos != (V{3, -7, -4}) ||
		moons[2].pos != (V{1, -7, 5}) ||
		moons[3].pos != (V{2, 2, 0}) {
		t.Fatalf("step didn't work: %v", moons)
	}
	if moons[0].vel != (V{3, -1, -1}) ||
		moons[1].vel != (V{1, 3, 3}) ||
		moons[2].vel != (V{-3, 1, -3}) ||
		moons[3].vel != (V{-1, -3, 1}) {
		t.Fatalf("step didn't work: %v", moons)
	}
	//After 1 step:
	//pos=<x= 2, y=-1, z= 1>, vel=<x= 3, y=-1, z=-1>
	//pos=<x= 3, y=-7, z=-4>, vel=<x= 1, y= 3, z= 3>
	//pos=<x= 1, y=-7, z= 5>, vel=<x=-3, y= 1, z=-3>
	//pos=<x= 2, y= 2, z= 0>, vel=<x=-1, y=-3, z= 1>

	Step(moons)
	if moons[0].pos != (V{5, -3, -1}) ||
		moons[1].pos != (V{1, -2, 2}) ||
		moons[2].pos != (V{1, -4, -1}) ||
		moons[3].pos != (V{1, -4, 2}) {
		t.Fatalf("step didn't work: %v", moons)
	}
	if moons[0].vel != (V{3, -2, -2}) ||
		moons[1].vel != (V{-2, 5, 6}) ||
		moons[2].vel != (V{0, 3, -6}) ||
		moons[3].vel != (V{-1, -6, 2}) {
		t.Fatalf("step didn't work: %v", moons)
	}
	//After 2 steps:
	//pos=<x= 5, y=-3, z=-1>, vel=<x= 3, y=-2, z=-2>
	//pos=<x= 1, y=-2, z= 2>, vel=<x=-2, y= 5, z= 6>
	//pos=<x= 1, y=-4, z=-1>, vel=<x= 0, y= 3, z=-6>
	//pos=<x= 1, y=-4, z= 2>, vel=<x=-1, y=-6, z= 2>
	moons = []*Moon{
		NewMoon(-8, -10, 0),
		NewMoon(5, 5, 10),
		NewMoon(2, -7, 3),
		NewMoon(9, -8, -3),
	}
	for i := 0; i < 100; i++ {
		Step(moons)
	}
	if moons[0].pos != (V{8, -12, -9}) ||
		moons[1].pos != (V{13, 16, -3}) ||
		moons[2].pos != (V{-29, -11, -1}) ||
		moons[3].pos != (V{16, -13, 23}) {
		t.Fatalf("step didn't work: %v", moons)
	}
	if moons[0].vel != (V{-7, 3, 0}) ||
		moons[1].vel != (V{3, -11, -5}) ||
		moons[2].vel != (V{-3, 7, 4}) ||
		moons[3].vel != (V{7, 1, 1}) {
		t.Fatalf("step didn't work: %v", moons)
	}

	if Energy(moons) != 1940 {
		t.Fatal("energy is wrong")
	}
	//pos=<x=  8, y=-12, z= -9>, vel=<x= -7, y=  3, z=  0>
	//pos=<x= 13, y= 16, z= -3>, vel=<x=  3, y=-11, z= -5>
	//pos=<x=-29, y=-11, z= -1>, vel=<x= -3, y=  7, z=  4>
	//pos=<x= 16, y=-13, z= 23>, vel=<x=  7, y=  1, z=  1>
}

func TestFindCycle(t *testing.T) {
	moons := []*Moon{
		NewMoon(-1, 0, 2),
		NewMoon(2, -10, -7),
		NewMoon(4, -8, 8),
		NewMoon(3, 5, -1),
	}

	if actual := FindCycle(moons); actual != 2772 {
		t.Fatalf("expected cycle at 2772, got %v", actual)
	}
}
