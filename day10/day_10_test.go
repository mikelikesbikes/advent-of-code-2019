package main

import (
	"math"
	"testing"
)

func TestFindBestAsteroid(t *testing.T) {
	input := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}

	actual, max := bestAsteroidCoord(parseInput(input))
	expected := Coord{x: 3, y: 4}
	if actual != expected && max != 8 {
		t.Fatalf("expected %v, got %v", expected, actual)
	}

	input = []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}

	actual, _ = bestAsteroidCoord(parseInput(input))
	expected = Coord{x: 11, y: 13}
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestAngle(t *testing.T) {
	if actual := (Coord{4, 4}).angle(Coord{4, 3}); actual != 0 {
		t.Fatalf("expected 0, got %v", actual)
	}

	if actual := (Coord{4, 4}).angle(Coord{5, 4}); actual != math.Pi/2 {
		t.Fatalf("expected %v, got %v", math.Pi/2, actual)
	}

	if actual := (Coord{4, 4}).angle(Coord{4, 5}); actual != math.Pi {
		t.Fatalf("expected %v, got %v", math.Pi, actual)
	}

	if actual := (Coord{4, 4}).angle(Coord{3, 4}); actual != 3*math.Pi/2 {
		t.Fatalf("expected %v, got %v", 3*math.Pi/2, actual)
	}
}
func TestFindNthAsteroid(t *testing.T) {
	input := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}
	start := Coord{3, 4}
	actual := findNthAsteroid(parseInput(input), start, 5)
	expected := Coord{4, 4}
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}

	input = []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}

	start = Coord{11, 13}
	actual = findNthAsteroid(parseInput(input), start, 200)
	expected = Coord{x: 8, y: 2}
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func Equal(a, b Coord) bool {
	return a.x == b.x && a.y == b.y
}
