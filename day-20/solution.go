// Can use hash table to track how many steps it takes to be a room
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
	// adjList := map[[2]int]map[[2]int]bool{}
	px, py := 0, 0
	for _, d := range sortedItems {
		choices := strings.Split(d, "|")
		for _, c := range choices {

			for i := 0; i < len(c); i++ {
				dx, dy := parseMove(c[i])
			}

		}
	}
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
