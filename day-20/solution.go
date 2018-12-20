// Still WIP, but thought the problem can be solved if correctly generate the adj list.
package main

import (
	"fmt"
	"math"
	"strings"
)

var (
	braceOpenASC  = "("[0]
	braceCloseASC = ")"[0]
	nASC          = "N"[0]
	sASC          = "S"[0]
	wASC          = "W"[0]
	eASC          = "E"[0]
)

type node struct {
	// When non-empty, is the leaf.
	content  string
	children []*node
}

func genBlocks(src string) []string {
	res := []string{}
	braceCounter, last := 0, 0
	for i := 0; i < len(src); i++ {
		if src[i] == braceOpenASC {
			if braceCounter == 0 {
				res, last = append(res, src[last:i]), i+1
			}
			braceCounter++
		}
		if src[i] == braceCloseASC {
			if braceCounter == 1 {
				res, last = append(res, src[last:i]), i+1
			}
			braceCounter--
		}
		if i == (len(src)-1) && last == 0 {
			res = append(res, src)
		}
	}
	return res
}

func genTree(src string, root *node) {
	if len(src) == 0 {
		return
	}
	blocks := genBlocks(src)
	if len(blocks) == 1 {
		child := &node{src, []*node{}}
		root.children = append(root.children, child)

		return
	}

	for _, b := range blocks {
		child := &node{"", []*node{}}
		root.children = append(root.children, child)
		genTree(b, child)
	}
}

func parseInput() string {
	s := ""
	fmt.Scan(&s)
	return s[1 : len(s)-1]
}

func walkTree(root *node, items []string) []string {
	if len(root.content) > 0 {
		items = append(items, root.content)
		return items
	}
	for _, c := range root.children {
		items = walkTree(c, items)
	}
	return items
}

func min(a, b int) int { return int(math.Min(float64(a), float64(b))) }
func max(a, b int) int { return int(math.Max(float64(a), float64(b))) }
func abs(a int) int    { return int(math.Abs(float64(a))) }

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
	root := &node{"", []*node{}}
	genTree(str, root)
	sortedItems := []string{}
	sortedItems = walkTree(root, sortedItems)
	fmt.Println(sortedItems, len(sortedItems))
	adjList := map[[2]int]map[[2]int]bool{}
	possiblePositions := map[[2]int]bool{[2]int{0, 0}: true}
	for _, d := range sortedItems {
		choices, newPositions := strings.Split(d, "|"), map[[2]int]bool{}
		fmt.Println(possiblePositions, "?")
		for p := range possiblePositions {
			for _, c := range choices {
				x, y := p[0], p[1]

				for i := 0; i < len(c); i++ {
					dx, dy := parseMove(c[i])
					if adjList[[2]int{x, y}] == nil {
						adjList[[2]int{x, y}] = map[[2]int]bool{}
					}
					adjList[[2]int{x, y}][[2]int{dx, dy}], x, y = true, x+dx, y+dy
					if adjList[[2]int{x, y}] == nil {
						adjList[[2]int{x, y}] = map[[2]int]bool{}
					}
					adjList[[2]int{x, y}][[2]int{-dx, -dy}] = true
				}

				newPositions[[2]int{x, y}] = true
			}
		}
		possiblePositions = newPositions
	}

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

func secondChallenge() {

}

func main() {
	str := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(str)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge()
}
