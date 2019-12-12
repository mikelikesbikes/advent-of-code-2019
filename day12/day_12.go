package main

import (
	"../utils"
	"fmt"
	"strings"
)

type V struct {
	x, y, z int
}

type Moon struct {
	pos, vel V
}

func NewMoon(x, y, z int) *Moon {
	return &Moon{pos: V{x, y, z}, vel: V{0, 0, 0}}
}

func ParseMoons(lines []string) []*Moon {
	var moons []*Moon
	for _, line := range lines {
		params := strings.Split(line[1:len(line)-1], ", ")
		moons = append(moons, NewMoon(
			utils.Atoi(params[0][2:]),
			utils.Atoi(params[1][2:]),
			utils.Atoi(params[2][2:])))
	}
	return moons
}

func delta(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

func (m *Moon) AdjustPos() {
	m.pos = V{m.pos.x + m.vel.x, m.pos.y + m.vel.y, m.pos.z + m.vel.z}
}

func (v V) Energy() int {
	return utils.Abs(v.x) + utils.Abs(v.y) + utils.Abs(v.z)
}
func (m *Moon) Energy() int {
	return m.pos.Energy() * m.vel.Energy()
}

func AdjustVelocity(m1, m2 *Moon) {
	dx := delta(m1.pos.x, m2.pos.x)
	dy := delta(m1.pos.y, m2.pos.y)
	dz := delta(m1.pos.z, m2.pos.z)
	m1.vel = V{m1.vel.x + dx, m1.vel.y + dy, m1.vel.z + dz}
	m2.vel = V{m2.vel.x - dx, m2.vel.y - dy, m2.vel.z - dz}
}

func Step(moons []*Moon) {
	AdjustVelocity(moons[0], moons[1])
	AdjustVelocity(moons[0], moons[2])
	AdjustVelocity(moons[0], moons[3])
	AdjustVelocity(moons[1], moons[2])
	AdjustVelocity(moons[1], moons[3])
	AdjustVelocity(moons[2], moons[3])

	for _, moon := range moons {
		moon.AdjustPos()
	}
}

func Energy(moons []*Moon) int {
	var total int
	for _, moon := range moons {
		total += moon.Energy()
	}
	return total
}

func FlattenX(moons []*Moon) string {
	var f []string
	for _, m := range moons {
		f = append(f, fmt.Sprintf("%v,%v", m.pos.x, m.vel.x))
	}
	return strings.Join(f, ":")
}
func FlattenY(moons []*Moon) string {
	var f []string
	for _, m := range moons {
		f = append(f, fmt.Sprintf("%v,%v", m.pos.y, m.vel.y))
	}
	return strings.Join(f, ":")
}
func FlattenZ(moons []*Moon) string {
	var f []string
	for _, m := range moons {
		f = append(f, fmt.Sprintf("%v,%v", m.pos.z, m.vel.z))
	}
	return strings.Join(f, ":")
}

func FindCycle(moons []*Moon) int {
	xs := make(map[string]bool)
	ys := make(map[string]bool)
	zs := make(map[string]bool)
	for {
		xk := FlattenX(moons)
		yk := FlattenY(moons)
		zk := FlattenZ(moons)
		if xs[xk] && ys[yk] && zs[zk] {
			break
		}
		xs[xk] = true
		ys[yk] = true
		zs[zk] = true
		Step(moons)
	}
	return utils.LCM(len(xs), len(ys), len(zs))
}

func main() {
	lines := utils.ReadLines("input.txt")
	moons := ParseMoons(lines)
	for i := 0; i < 1000; i++ {
		Step(moons)
	}
	fmt.Println(Energy(moons))

	moons = ParseMoons(lines)
	fmt.Println(FindCycle(moons))
}
