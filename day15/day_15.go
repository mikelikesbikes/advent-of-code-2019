package main

import (
	"../utils"
	"../utils/intcode"
	"fmt"
	"time"
)

type XY struct {
	x, y int
}

type RepairDroid struct {
	canvas   map[XY]int
	pos      XY
	target   XY
	computer *intcode.Program
	path     []XY
}

func (a XY) adjacents() []XY {
	return []XY{
		{a.x + 1, a.y},
		{a.x, a.y + 1},
		{a.x - 1, a.y},
		{a.x, a.y - 1},
	}
}

func dir(a, b XY) int {
	dx := b.x - a.x
	dy := b.y - a.y
	var res int
	if dx == 0 && dy == -1 {
		res = 1
	} else if dx == 0 && dy == 1 {
		res = 2
	} else if dx == -1 && dy == 0 {
		res = 3
	} else if dx == 1 && dy == 0 {
		res = 4
	} else {
		panic(fmt.Sprintf("%v and %v are NOT adjacent", a, b))
	}
	return res
}

func NewRepairDroid(code string) *RepairDroid {
	return &RepairDroid{canvas: make(map[XY]int), computer: intcode.ParseIntcodeProgram(code), pos: XY{0, 0}, path: make([]XY, 0)}
}

func (rd *RepairDroid) availablePos() []XY {
	var available []XY
	for _, pos := range rd.pos.adjacents() {
		if _, visited := rd.canvas[pos]; !visited {
			available = append(available, pos)
		}
	}
	return available
}

func findBounds(m map[XY]int) (XY, XY) {
	var minx, miny, maxx, maxy int
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

func (rd *RepairDroid) Render() {
	fmt.Println("\033[2J")

	min, max := XY{-21, -21}, XY{19, 19}
	for y := min.y; y <= max.y; y++ {
		fmt.Print("\n")
		for x := min.x; x <= max.x; x++ {
			if rd.pos.x == x && rd.pos.y == y {
				fmt.Print("D")
			} else {
				v, visited := rd.canvas[XY{x, y}]
				if visited {
					switch v {
					case 0:
						fmt.Print("#")
					case 1:
						fmt.Print(".")
					case 2:
						fmt.Print("O")
					}
				} else {
					fmt.Print(" ")
				}
			}
		}
	}
	fmt.Println()
	time.Sleep(15 * time.Millisecond)
}

func (rd *RepairDroid) nextMove() int {
	available := rd.availablePos()
	if len(available) == 0 {
		last := len(rd.path) - 1
		rd.target = rd.path[last]
		rd.path = rd.path[:last]
	} else {
		rd.target = available[0]
	}
	dir := dir(rd.pos, rd.target)
	return dir
}

func (rd *RepairDroid) move(status int) {
	_, path := rd.canvas[rd.target]
	rd.canvas[rd.target] = status
	if status != 0 {
		if !path {
			rd.path = append(rd.path, rd.pos)
		}
		rd.pos = rd.target
	}
}

func distMap(m map[XY]int, src XY) map[XY]int {
	visited := make(map[XY]int, len(m))
	visited[src] = 0
	queue := []XY{src}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		dist := visited[first]
		for _, pos := range first.adjacents() {
			if _, v := visited[pos]; !v {
				if v, found := m[pos]; found && v != 0 {
					visited[pos] = dist + 1
					queue = append(queue, pos)
				}
			}
		}
	}
	return visited
}

func (rd *RepairDroid) Start(render bool) {
	go func() {
		rd.computer.Run()
	}()

	for !(len(rd.path) == 0 && len(rd.availablePos()) == 0) {
		select {
		case status := <-rd.computer.Out:
			rd.move(status)
			if render {
				rd.Render()
			}
		case <-rd.computer.InputNeeded:
			rd.computer.In <- rd.nextMove()
		}
	}
}

func main() {
	lines := utils.ReadLines("input.txt")
	droid := NewRepairDroid(lines[0])

	droid.Start(false)

	var oxygen XY
	for k, v := range droid.canvas {
		if v == 2 {
			oxygen = k
			break
		}
	}
	m := distMap(droid.canvas, XY{0, 0})
	fmt.Println(m[oxygen])

	m2 := distMap(droid.canvas, oxygen)
	var maxLen int
	for _, v := range m2 {
		if v > maxLen {
			maxLen = v
		}
	}
	fmt.Println(maxLen)
}
