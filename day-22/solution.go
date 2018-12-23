// For question 2, Dijkstra's algorithm would be much faster but require more lines of code.
package main

import (
	"fmt"
	"math"
)

const (
	M              = 20183
	torch          = 0
	climbGear      = 1
	none           = 2
	switchGearTime = 7
)

type edge struct {
	from, to [2]int
	cost     int
}

var dirs = [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
var allowedGearsMap = map[int][]int{
	0: []int{torch, climbGear},
	1: []int{climbGear, none},
	2: []int{torch, none},
}

func genMap(d, w, h int) ([][]int, int, int) {
	extend := getExtend(w, h)
	e := make([][]int, h+extend)
	for i := 0; i < h+extend; i++ {
		e[i] = make([]int, w+extend)
	}

	e[0][0], e[h-1][w-1] = d, d
	for i := 0; i < (h - 1 + extend); i++ {
		e[i+1][0] = ((i+1)*48271 + d) % M
	}
	for i := 0; i < (w - 1 + extend); i++ {
		e[0][i+1] = ((i+1)*16807 + d) % M
	}
	for i := 1; i < h+extend; i++ {
		for j := 1; j < w+extend; j++ {
			// Don't have to re-calc the ending point
			if (i == (h - 1)) && (j == (w - 1)) {
				continue
			}
			e[i][j] = (e[i-1][j]*e[i][j-1] + d) % M
		}
	}
	for i := 0; i < h+extend; i++ {
		for j := 0; j < w+extend; j++ {
			e[i][j] = e[i][j] % 3
		}
	}
	return e, h + extend, w + extend
}

func getAnotherGear(gears []int, gear int) int {
	for _, r := range gears {
		if r != gear {
			return r
		}
	}
	return -1
}

func getExtend(w, h int) int {
	return (w + h) / 2
}

func firstChallenge(d, w, h int) {
	m, _, _ := genMap(d, w, h)
	sum := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			sum += m[i][j]
		}
	}
	fmt.Println(sum)
}

func secondChallenge(d, w, h int) {
	m, mapH, mapW := genMap(d, w, h)
	distance, cannotReach := map[[2]int][]int{}, math.MaxInt32
	for i := 0; i < mapH; i++ {
		for j := 0; j < mapW; j++ {
			distance[[2]int{i, j}] = []int{torch: cannotReach, climbGear: cannotReach, none: cannotReach}
		}
	}
	distance[[2]int{0, 0}][torch] = 0 // The starting point is rocky.

	visit := func(i, j int) bool {
		changed := false
		for _, r := range allowedGearsMap[m[i][j]] {
			another := getAnotherGear(allowedGearsMap[m[i][j]], r)
			newDistance := distance[[2]int{i, j}][r] + switchGearTime
			if newDistance < distance[[2]int{i, j}][another] {
				distance[[2]int{i, j}][another], changed = newDistance, true
			}
		}
		for _, dir := range dirs {
			x, y := j+dir[1], i+dir[0]
			if (x < 0) || (x == mapW) || (y < 0) || (y == mapH) {
				continue
			}
			for _, r := range allowedGearsMap[m[y][x]] {
				newDistance := distance[[2]int{i, j}][r] + 1
				if newDistance < distance[[2]int{y, x}][r] {
					distance[[2]int{y, x}][r], changed = newDistance, true
				}
			}
		}
		return changed
	}

	for n := 0; n < w+h; n++ {
		done := true
		for i := 0; i < mapH; i++ {
			for j := 0; j < mapW; j++ {
				if visit(i, j) {
					done = false
				}
			}
		}
		if done {
			break
		}
	}
	fmt.Println(distance[[2]int{h, w}][torch])
}

func parseInput() (int, int, int) {
	d, w, h := 0, 0, 0
	fmt.Scanf("depth: %d", &d)
	fmt.Scanf("target: %d,%d", &w, &h)
	return d, w + 1, h + 1
}

func main() {
	d, w, h := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(d, w, h)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(d, w, h)
}
