package main

import (
	"../utils"
	"../utils/intcode"
	"fmt"
)

type XY struct {
	x, y int
}

type Beam struct {
	computer *intcode.Program
	picture  map[XY]int
	code     string
}

func NewBeam(code string) *Beam {
	return &Beam{computer: intcode.ParseIntcodeProgram(code), picture: make(map[XY]int, 0), code: code}
}

func (b *Beam) Start(pos XY) {
	go b.computer.Run()
	var toggle bool
loop:
	for {
		select {
		case <-b.computer.InputNeeded:
			if toggle {
				b.computer.In <- pos.y
			} else {
				b.computer.In <- pos.x
			}
			toggle = !toggle
		case c := <-b.computer.Out:
			b.picture[pos] = c
			break loop
		}
	}
}

func (b *Beam) TestPos(pos XY) int {
	computer := intcode.ParseIntcodeProgram(b.code)
	go computer.Run()
	var toggle bool
	for {
		select {
		case <-computer.InputNeeded:
			if toggle {
				computer.In <- pos.y
			} else {
				computer.In <- pos.x
			}
			toggle = !toggle
		case c := <-computer.Out:
			return c
		}
	}
}

func (b *Beam) TestField(minx, miny, maxx, maxy int) {
	for y := miny; y < maxy; y++ {
		for x := minx; x < maxx; x++ {
			pos := XY{x, y}
			b.picture[pos] = b.TestPos(pos)
		}
	}
}
func findBounds(m map[XY]int) (XY, XY) {
	var minx, miny, maxx, maxy int
	for k, _ := range m {
		minx = k.x
		maxx = k.x
		miny = k.y
		maxy = k.y
		break
	}
	for p := range m {
		if p.x < minx {
			minx = p.x
		}
		if p.y < miny {
			miny = p.y
		}
		if p.x > maxx {
			maxx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
	}
	return XY{minx, miny}, XY{maxx, maxy}
}
func Render(m map[XY]int) {
	min, max := findBounds(m)
	fmt.Println(min, max)
	fmt.Print("       ")
	for x := min.x; x <= max.x; x++ {
		fmt.Print((x / 10) % 10)
	}
	fmt.Print("\n       ")
	for x := min.x; x <= max.x; x++ {
		fmt.Print(x % 10)
	}
	for y := min.y; y <= max.y; y++ {
		fmt.Printf("\n%06d ", y)
		for x := min.x; x <= max.x; x++ {
			switch m[XY{x, y}] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			}
		}
	}
	fmt.Println()
}

func main() {
	lines := utils.ReadLines("input.txt")

	beam := NewBeam(lines[0])
	beam.TestField(0, 0, 50, 50)

	var affected int
	for _, v := range beam.picture {
		if v == 1 {
			affected++
		}
	}

	fmt.Println(affected)

	x, y := 80, 89
	for w := 8; w < 100; w += 2 {
		if (w/2)%2 == 0 {
			x += 25
			y += 28
		} else {
			x += 20
			y += 22
		}
	}

	//beam = NewBeam(lines[0])
	//beam.TestField(x, y, x+120, y+120)
	//Render(beam.picture)
	// this got me close... but it's off by 1... i think... so

	// 1122,1248
	beam = NewBeam(lines[0])
	beam.TestField(1122, 1248, 1222, 1348)
	Render(beam.picture)

	fmt.Println(x*10000 + y)
}
