package main

import (
	"fmt"
	"math"
	"os"
)

const (
	max    = 1000
	theDig = 100
	X      = 300
	Y      = 300
)

func getScore(x, y, s int) int {
	r := x + 10
	p := (r*y + s) % max
	p = p * r % max
	return p/theDig - 5
}

func getInput() int {
	s := 0
	fmt.Scan(&s)
	return s
}

func getGroupScore(m [][]int, x, y, z int) int {
	sum := 0
	for i := 0; i < z; i++ {
		for j := 0; j < z; j++ {
			sum += m[y+i][x+j]
		}
	}
	return sum
}

func getBestForSquareZ(m [][]int, s, z int) ([2]int, int) {
	best, max := [2]int{-1, -1}, math.MinInt32
	for i := 0; i < (Y - z + 1); i++ {
		for j := 0; j < (X - z + 1); j++ {
			gS := getGroupScore(m, j, i, z)
			if gS > max {
				max, best = gS, [2]int{j + 1, i + 1}
			}
		}
	}
	return best, max
}

func firstChallenge() {
	s := getInput()
	m := make([][]int, Y)
	for i := 0; i < Y; i++ {
		m[i] = make([]int, X)
	}
	for i := 0; i < Y; i++ {
		for j := 0; j < X; j++ {
			m[i][j] = getScore(j+1, i+1, s)
		}
	}
	best, _ := getBestForSquareZ(m, s, 3)
	fmt.Println(best)
}

func secondChallenge() {
	s := getInput()
	m := make([][]int, Y)
	for i := 0; i < Y; i++ {
		m[i] = make([]int, X)
	}
	for i := 0; i < Y; i++ {
		for j := 0; j < X; j++ {
			m[i][j] = getScore(j+1, i+1, s)
		}
	}
	best, max := [3]int{-1, -1, -1}, math.MinInt32
	for i := 1; i <= 300; i++ {
		xy, score := getBestForSquareZ(m, s, i)
		if score > max {
			best, max = [3]int{xy[0], xy[1], i}, score
		}
	}

	fmt.Println(best)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
