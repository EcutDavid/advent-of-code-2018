// Using a doubly linked list to sort all the items, loop them until there is a loop has no collision,
// whenever there is a collision, drop two nodes from the list.
// Using a linked list instead of an array so that we don't spend time on items that already "consumed".
package main

import (
	"fmt"
	"math"
)

var uAAsc = "A"[0]
var uZAsc = "Z"[0]
var lAAsc = "a"[0]

type node struct {
	val  int
	next *node
	prev *node
}

func parseInput() []int {
	src := ""
	fmt.Scan(&src)
	dst := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] >= uAAsc && src[i] <= uZAsc {
			dst[i] = int(src[i]) - int(uAAsc) + 1
		} else {
			dst[i] = -(int(src[i]) - int(lAAsc) + 1)
		}
	}
	return dst
}

func genRootNode(a []int, f1, f2 int) *node {
	b := make([]int, 0, len(a))
	for _, v := range a {
		if v != f1 && v != f2 {
			b = append(b, v)
		}
	}
	root := &node{b[0], nil, nil}
	prev := root
	for i := 1; i < len(b); i++ {
		cur := &node{b[i], nil, nil}
		prev.next, cur.prev, prev = cur, prev, cur
	}
	return root
}

func canCollide(a *node) bool {
	if (a.next != nil) && (a.val == -a.next.val) {
		return true
	}
	return false
}

// Root may change, so returns the new one.
func simulate(root *node) *node {
	for {
		next, hasCollide := root, false
		for next != nil {
			if !canCollide(next) {
				next = next.next
				continue
			}
			hasCollide = true
			if next.prev == nil { // Only root has no prev.
				root = next.next.next
				if root != nil {
					root.prev = nil
				}
				break
			}
			next.prev.next = next.next.next
			if next.next.next != nil {
				next.next.next.prev = next.prev
			}
			next = next.next.next
		}
		if !hasCollide {
			break
		}
	}
	return root
}

func getLength(root *node) int {
	length, next := 0, root
	for next != nil {
		length, next = length+1, next.next
	}
	return length
}

func firstChallenge(a []int) {
	// -100 just a random number choosed that bigger than 27
	root := genRootNode(a, -100, -100)
	root = simulate(root)
	fmt.Println(getLength(root))
}

func secondChallenge(a []int) {
	min := math.MaxInt32
	for i := 1; i < 27; i++ {
		root := genRootNode(a, -i, i)
		root = simulate(root)
		choice := getLength(root)
		if choice < min {
			min = choice
		}
	}
	fmt.Println(min)
}

func main() {
	a := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(a)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(a)
}
