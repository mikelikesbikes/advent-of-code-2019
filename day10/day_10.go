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

func (c1 Coord) angle(c2 Coord) float64 {
	dx := c2.x - c1.x
	dy := c2.y - c1.y
	angle := (math.Atan2(float64(dy), float64(dx)) + math.Pi/2)
	if angle < 0 {
		angle = angle + 2*math.Pi
	}
	return angle
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

func mapAsteroidAngles(asteroids []Coord, coord Coord) map[float64][]Coord {
	am := make(map[float64][]Coord)
	for _, a := range asteroids {
		if a != coord {
			angle := coord.angle(a)
			am[angle] = append(am[angle], a)
		}
	}
	return am
}

func bestAsteroidCoord(asteroids []Coord) (Coord, int) {
	var maxAsteroid Coord
	var max int
	for _, a := range asteroids {
		m := mapAsteroidAngles(asteroids, a)
		if len(m) > max {
			maxAsteroid = a
			max = len(m)
		}
	}
	return maxAsteroid, max
}

func findNthAsteroid(asteroids []Coord, coord Coord, n int) Coord {
	am := mapAsteroidAngles(asteroids, coord)
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
