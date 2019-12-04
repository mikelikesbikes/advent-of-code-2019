package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type Blaze struct {
	log map[string]int
}

func NewBlaze() *Blaze {
	return &Blaze{log: make(map[string]int)}
}

type Heading struct {
	dir string
	len int
}

func NewHeading(s string) Heading {
	len, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(fmt.Sprintf("can't convert %v to an int", s[1:]))
	}
	return Heading{dir: string(s[0]), len: len}
}

func (blaze *Blaze) Mark(marker string, i int) {
	_, ok := blaze.log[marker]
	if ok {
		return
	}
	blaze.log[marker] = i
}

func (pos *Pos) Inc(diff []int) *Pos {
	return &Pos{x: pos.x + diff[0], y: pos.y + diff[1]}
}

func (pos *Pos) key() string {
	return fmt.Sprintf("%v,%v", pos.x, pos.y)
}

func dist(pos Pos) int {
	x := pos.x
	if x < 0 {
		x = -x
	}
	y := pos.y
	if y < 0 {
		y = -y
	}
	return x + y
}

func DistanceToCross(w1, w2 string) int {
	pathMap := make(map[Pos]*Blaze)
	buildPath(w1, pathMap, "1")
	buildPath(w2, pathMap, "2")
	dmin := 9223372036854775807
	for pos, v := range pathMap {
		if len(v.log) > 1 {
			d := dist(pos)
			if d < dmin {
				dmin = d
			}
		}
	}
	return dmin
}

func EarliestCross(w1, w2 string) int {
	pathMap := make(map[Pos]*Blaze)
	buildPath(w1, pathMap, "1")
	buildPath(w2, pathMap, "2")
	smin := 9223372036854775807
	for _, v := range pathMap {
		if len(v.log) > 1 {
			steps := 0
			for _, s := range v.log {
				steps += s
			}
			if steps < smin {
				smin = steps
			}
		}
	}
	return smin
}

func buildPath(ps string, m map[Pos]*Blaze, marker string) {
	var headings []Heading
	for _, s := range strings.Split(ps, ",") {
		headings = append(headings, NewHeading(s))
	}

	pos := &Pos{0, 0}
	steps := 0
	for _, heading := range headings {
		var diff []int
		switch heading.dir {
		case "R":
			diff = []int{1, 0}
		case "L":
			diff = []int{-1, 0}
		case "U":
			diff = []int{0, -1}
		case "D":
			diff = []int{0, 1}
		}
		for i := 0; i < heading.len; i++ {
			pos = pos.Inc(diff)
			steps += 1
			if _, ok := m[*pos]; !ok {
				m[*pos] = NewBlaze()
			}
			m[*pos].Mark(marker, steps)
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open input.txt")
	}

	scanner := bufio.NewScanner(file)
	var wires []string
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}

	fmt.Println(DistanceToCross(wires[0], wires[1]))
	fmt.Println(EarliestCross(wires[0], wires[1]))
}
