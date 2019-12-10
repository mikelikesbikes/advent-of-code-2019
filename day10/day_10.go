package main

import (
	"../utils"
	"fmt"
	"math"
	"sort"
)

// map polar coords from each asteriod to all other asteroids storing only the angle

type Coord struct {
	x, y int
}

func abs(a int) int {
	aa := a
	if aa < 0 {
		aa = -a
	}
	return aa
}

func gcd(a, b int) int {
	aa := abs(a)
	ab := abs(b)
	sm := aa
	if sm > ab {
		sm = ab
	}

	if aa == 0 {
		return ab
	} else if ab == 0 {
		return aa
	}

	for ; sm > 0; sm-- {
		if aa%sm == 0 && ab%sm == 0 {
			return sm
		}
	}
	return 1
}

func ratio(x, y int) Coord {
	gcd := gcd(x, y)
	return Coord{(x / gcd), (y / gcd)}
}

func (c1 Coord) angle(c2 Coord) Coord {
	dx := c2.x - c1.x
	dy := c2.y - c1.y
	return ratio(dx, dy)
}

func parseInput(input []string) []Coord {
	var asteroids []Coord
	for y := 0; y < len(input); y++ {
		for x, a := range input[y] {
			if a == '#' {
				asteroids = append(asteroids, Coord{x, y})
			}
		}
	}
	return asteroids
}

func bestAsteroidCoord(asteroids []Coord) (Coord, int) {
	var maxAsteroid Coord
	var max int
	for i, a1 := range asteroids {

		m := make(map[Coord]bool)
		for j, a2 := range asteroids {
			if i == j {
				continue
			}
			angle := a1.angle(a2)
			m[angle] = true
		}

		if len(m) > max {
			maxAsteroid = a1
			max = len(m)
		}
	}

	return maxAsteroid, max
}

func (c1 Coord) angle2(c2 Coord) float64 {
	dx := c2.x - c1.x
	dy := c2.y - c1.y
	angle := (math.Atan2(float64(dy), float64(dx)) + math.Pi/2)
	if angle < 0 {
		angle = angle + 2*math.Pi
	}
	return angle
}

func findNthAsteroid(asteroids []Coord, coord Coord, n int) Coord {
	am := make(map[float64][]Coord)
	for _, a := range asteroids {
		if a.x == coord.x && a.y == coord.y {
			continue
		}
		angle := coord.angle2(a)
		am[angle] = append(am[angle], a)
	}
	keys := make([]float64, 0, len(am))
	for k := range am {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	return am[keys[n-1]][0]
}

func main() {
	lines := utils.ReadLines("input.txt")
	asteroids := parseInput(lines)

	best, max := bestAsteroidCoord(asteroids)
	fmt.Println(max)

	coord := findNthAsteroid(asteroids, best, 200)
	fmt.Println(coord.x*100 + coord.y)
}
