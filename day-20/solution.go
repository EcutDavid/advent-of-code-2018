// More comments to be done, a good one!
package main

import (
	"fmt"
)

var (
	braceOpenASC  = "("[0]
	braceCloseASC = ")"[0]
	nASC          = "N"[0]
	sASC          = "S"[0]
	wASC          = "W"[0]
	eASC          = "E"[0]
	divASC        = "|"[0]
)

func findBraceCloseIndex(src string, index int) int {
	braceCounter := 0
	for i := index + 1; i < len(src); i++ {
		if src[i] == braceOpenASC {
			braceCounter++
		}
		if src[i] == braceCloseASC {
			if braceCounter == 0 {
				return i
			}
			braceCounter--
		}
	}
	return -1
}

func genChoices(src string) []string {
	res := []string{}
	braceCounter, last := 0, 0
	for i := 0; i < len(src); i++ {
		if src[i] == braceOpenASC {
			braceCounter++
		}
		if src[i] == braceCloseASC {
			braceCounter--
		}
		if src[i] == divASC && braceCounter == 0 {
			res, last = append(res, src[last:i]), i+1
		}
		if i == (len(src) - 1) {
			res = append(res, src[last:])
		}
	}
	return res
}

func parseInput() string {
	s := ""
	fmt.Scan(&s)
	return s[1 : len(s)-1]
}

var adjList = map[[2]int]map[[2]int]bool{}

func walk(src string, positions [][2]int) [][2]int {
	if len(src) == 0 {
		return positions
	}
	// Otherwise, walk choice gonna mutate positions, which is unexpected.
	positionsCopy := make([][2]int, len(positions))
	copy(positionsCopy, positions)
	positions = positionsCopy
	for i := 0; i < len(src); i++ {
		if src[i] == braceOpenASC {
			newPositionsMap, closeIndex := map[[2]int]bool{}, findBraceCloseIndex(src, i)
			cs := genChoices(src[i+1 : closeIndex])
			for _, c := range cs {
				newPositions := walk(c, positions)
				for _, p := range newPositions {
					newPositionsMap[p] = true
				}
			}
			newPositions := [][2]int{}
			for k := range newPositionsMap {
				newPositions = append(newPositions, k)
			}
			walk(src[closeIndex+1:], newPositions)
			return newPositions
		}
		dx, dy := parseMove(src[i])
		for k, p := range positions {
			x, y := p[0], p[1]
			if adjList[[2]int{x, y}] == nil {
				adjList[[2]int{x, y}] = map[[2]int]bool{}
			}
			adjList[[2]int{x, y}][[2]int{dx, dy}], x, y = true, x+dx, y+dy
			if adjList[[2]int{x, y}] == nil {
				adjList[[2]int{x, y}] = map[[2]int]bool{}
			}
			adjList[[2]int{x, y}][[2]int{-dx, -dy}] = true
			positions[k] = [2]int{x, y}
		}
	}
	return positions
}

func parseMove(move byte) (int, int) {
	x, y := 0, 0
	switch move {
	case nASC:
		y++
	case sASC:
		y--
	case wASC:
		x++
	case eASC:
		x--
	}
	return x, y
}

func firstChallenge(str string) {
	adjList = map[[2]int]map[[2]int]bool{}
	walk(str, [][2]int{[2]int{0, 0}})

	visited, distance := map[[2]int]bool{[2]int{0, 0}: true}, map[[2]int]int{[2]int{0, 0}: 0}
	queue, maxDistance := [][2]int{[2]int{0, 0}}, 0
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]
		if maxDistance < distance[task] {
			maxDistance = distance[task]
		}
		for p := range adjList[task] {
			room := [2]int{task[0] + p[0], task[1] + p[1]}
			if visited[room] {
				continue
			}
			queue, distance[room], visited[room] = append(queue, room), distance[task]+1, true
		}
	}
	fmt.Println(maxDistance)
}

func secondChallenge(str string) {
	adjList = map[[2]int]map[[2]int]bool{}
	walk(str, [][2]int{[2]int{0, 0}})

	visited, distance := map[[2]int]bool{[2]int{0, 0}: true}, map[[2]int]int{[2]int{0, 0}: 0}
	queue, sum := [][2]int{[2]int{0, 0}}, 0
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]
		if distance[task] >= 1000 {
			sum++
		}
		for p := range adjList[task] {
			room := [2]int{task[0] + p[0], task[1] + p[1]}
			if visited[room] {
				continue
			}
			queue, distance[room], visited[room] = append(queue, room), distance[task]+1, true
		}
	}
	fmt.Println(sum)
}

func main() {
	str := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(str)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(str)
}
