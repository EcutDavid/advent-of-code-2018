// Parsing a graph from the input, then, DFS to walk through the map.
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
var vectorMap = map[byte][2]int{
	nASC: [2]int{0, 1},
	sASC: [2]int{0, -1},
	wASC: [2]int{-1, 0},
	eASC: [2]int{1, 0},
}

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

func setAdjList(adjList map[[2]int]map[[2]int]bool, pos, vector [2]int) {
	if adjList[pos] == nil {
		adjList[pos] = map[[2]int]bool{}
	}
	adjList[pos][vector] = true
}

func walk(route string, positions [][2]int, adjList map[[2]int]map[[2]int]bool) [][2]int {
	if len(route) == 0 {
		return positions
	}

	for i := 0; i < len(route); i++ {
		if route[i] == braceOpenASC {
			positionsSet, closeIndex := map[[2]int]bool{}, findBraceCloseIndex(route, i)
			choices, newPositions := genChoices(route[i+1:closeIndex]), [][2]int{}
			for _, c := range choices {
				// Copies the positions so that positions will remain unchanged after walking a choice.
				positionsCopy := make([][2]int, len(positions))
				copy(positionsCopy, positions)
				for _, p := range walk(c, positionsCopy, adjList) {
					if !positionsSet[p] {
						newPositions = append(newPositions, p)
					}
					positionsSet[p] = true
				}
			}
			positions, i = newPositions, closeIndex
			continue
		}

		v := vectorMap[route[i]]
		for k, p := range positions {
			setAdjList(adjList, p, v)
			newPos := [2]int{p[0] + v[0], p[1] + v[1]}
			setAdjList(adjList, newPos, [2]int{-v[0], -v[1]})
			positions[k] = newPos
		}
	}
	return positions
}

func bfs(adjList map[[2]int]map[[2]int]bool, visit func(distance int)) {
	visited, distance := map[[2]int]bool{[2]int{0, 0}: true}, map[[2]int]int{[2]int{0, 0}: 0}
	queue := [][2]int{[2]int{0, 0}}
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]
		visit(distance[task])
		for p := range adjList[task] {
			room := [2]int{task[0] + p[0], task[1] + p[1]}
			if visited[room] {
				continue
			}
			queue, distance[room], visited[room] = append(queue, room), distance[task]+1, true
		}
	}
}

func firstChallenge(str string) {
	adjList := map[[2]int]map[[2]int]bool{}
	walk(str, [][2]int{[2]int{0, 0}}, adjList)

	maxDistance := 0
	visit := func(distance int) {
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	bfs(adjList, visit)
	fmt.Println(maxDistance)
}

func secondChallenge(str string) {
	adjList := map[[2]int]map[[2]int]bool{}
	walk(str, [][2]int{[2]int{0, 0}}, adjList)

	sum := 0
	visit := func(distance int) {
		if distance >= 1000 {
			sum++
		}
	}
	bfs(adjList, visit)
	fmt.Println(sum)
}

func parseInput() string {
	s := ""
	fmt.Scan(&s)
	return s[1 : len(s)-1]
}

func main() {
	str := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(str)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(str)
}
