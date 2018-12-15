// This one is so hard(for the current me)... Sorry for the messy code, gonna spend 0.5 hour this weekend clean the code.
// Good problem!
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	left  = [2]int{0, -1}
	right = [2]int{0, 1}
	up    = [2]int{-1, 0}
	down  = [2]int{1, 0}
	dirs  = [4][2]int{up, right, down, left}
)

var (
	forwardSlash  = "/"[0]
	backwardSlash = "\\"[0]
	empty         = " "[0]
	plus          = "+"[0]
	byteUp        = "^"[0]
	byteRight     = ">"[0]
	byteDown      = "v"[0]
	byteLeft      = "<"[0]
)

type cart struct {
	pos         [2]int
	dirPointer  int
	dirSwitcher int
	die         bool
}

func getInput() []string {
	lines := []string{}
	max := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) > max {
			max = len(scanner.Text())
		}
		lines = append(lines, scanner.Text())
	}
	for i := 0; i < len(lines); i++ {
		for j := len(lines[i]); j < max; j++ {
			lines[i] += " "
		}
	}
	return lines
}

func genTracks(lines []string) [][2][2]int {
	tracks, done := [][2][2]int{}, map[[2]int]bool{}
	W, H := len(lines[0]), len(lines)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if done[[2]int{i, j}] {
				continue
			}
			if lines[i][j] != forwardSlash {
				continue
			}
			w, h := 1, 1

			for lines[i][j+w] != backwardSlash {
				w++
			}
			for lines[i+h][j] != backwardSlash {
				h++
			}
			track := [2][2]int{
				[2]int{i, j}, [2]int{i + h, j + w},
			}
			tracks, done[track[1]] = append(tracks, track), true
		}
	}
	return tracks
}

func genAdjList(tracks [][2][2]int, lines []string) map[[2]int]map[[2]int]bool {
	adj := map[[2]int]map[[2]int]bool{}

	for _, t := range tracks {
		x1, y1, x2, y2 := t[0][1], t[0][0], t[1][1], t[1][0]
		adj[[2]int{y1, x1}], adj[[2]int{y1, x2}] = map[[2]int]bool{down: true, right: true}, map[[2]int]bool{down: true, left: true}
		adj[[2]int{y2, x1}], adj[[2]int{y2, x2}] = map[[2]int]bool{up: true, right: true}, map[[2]int]bool{up: true, left: true}

		for x := x1 + 1; x < x2; x++ {
			adj[[2]int{y1, x}], adj[[2]int{y2, x}] = map[[2]int]bool{left: true, right: true}, map[[2]int]bool{left: true, right: true}
		}
		for y := y1 + 1; y < y2; y++ {
			adj[[2]int{y, x1}], adj[[2]int{y, x2}] = map[[2]int]bool{up: true, down: true}, map[[2]int]bool{up: true, down: true}
		}
	}

	W, H := len(lines[0]), len(lines)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if lines[i][j] == plus {
				adj[[2]int{i, j}] = map[[2]int]bool{left: true, right: true, up: true, down: true}
			}
		}
	}
	return adj
}

type cartList []cart

func getCarts(lines []string) cartList {
	carts := []cart{}
	W, H := len(lines[0]), len(lines)
	cartMap := map[byte]int{byteUp: 0, byteRight: 1, byteDown: 2, byteLeft: 3}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			dir, ok := cartMap[lines[i][j]]
			if !ok {
				continue
			}
			carts = append(carts, cart{[2]int{i, j}, dir, 0, false})
		}
	}
	return cartList(carts)
}

func (c cartList) Len() int      { return len(c) }
func (c cartList) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c cartList) Less(i, j int) bool {
	if c[i].pos[0] == c[j].pos[0] {
		return c[i].pos[1] < c[j].pos[1]
	}
	return c[i].pos[0] < c[j].pos[0]
}
func (c *cart) move(adjList map[[2]int]map[[2]int]bool) [2]int {
	move := dirs[c.dirPointer]
	c.pos = [2]int{c.pos[0] + move[0], c.pos[1] + move[1]}
	newMoves := adjList[c.pos]
	if len(newMoves) == 4 {
		if c.dirSwitcher == 0 {
			c.dirPointer = (c.dirPointer + 4 - 1) % 4
		}
		if c.dirSwitcher == 2 {
			c.dirPointer = (c.dirPointer + 1) % 4
		}
		c.dirSwitcher = (c.dirSwitcher + 1) % 3
	} else {
		if !((newMoves[up] && newMoves[down]) || (newMoves[left] && newMoves[right])) {
			antiClock := true
			if newMoves[left] && newMoves[down] && move == right {
				antiClock = false
			}
			if newMoves[left] && newMoves[up] && move == down {
				antiClock = false
			}
			if newMoves[right] && newMoves[up] && move == left {
				antiClock = false
			}
			if newMoves[right] && newMoves[down] && move == up {
				antiClock = false
			}
			if antiClock {
				c.dirPointer = (c.dirPointer + 4 - 1) % 4
			} else {
				c.dirPointer = (c.dirPointer + 1) % 4
			}
		}
	}
	return c.pos
}

func simulate(adjList map[[2]int]map[[2]int]bool, carts cartList) {
	for {
		posMap := map[[2]int]bool{}
		// Have to sort to detect the first collide, if no need for detect first, we don't care about order.
		sort.Sort(carts)
		for i := 0; i < len(carts); i++ {
			pos := carts[i].move(adjList)
			if posMap[pos] {
				fmt.Printf("%d,%d\n", pos[1], pos[0])
				os.Exit(0)
			}
			posMap[pos] = true
		}
	}
}

func firstChallenge() {
	l := getInput()
	t := genTracks(l)
	adj := genAdjList(t, l)

	carts := getCarts(l)
	simulate(adj, carts)
}

func simulate2(adjList map[[2]int]map[[2]int]bool, carts cartList) {
	for {
		posMap := map[[2]int]int{}
		sum, theCart := 0, carts[0]
		for _, v := range carts {
			if !v.die {
				sum++
				theCart = v
			}
		}
		if sum == 1 {
			pos := theCart.pos
			fmt.Printf("%d,%d\n", pos[1], pos[0])
			os.Exit(0)
		}

		sort.Sort(carts)
		for i := 0; i < len(carts); i++ {
			if carts[i].die {
				continue
			}

			pos := carts[i].move(adjList)
			for k, v := range carts {
				if (k == i) || v.die {
					continue
				}
				if v.pos == pos {
					carts[k].die, carts[i].die = true, true
					break
				}
			}
			if carts[i].die {
				continue
			}
			if posMap[pos] != 0 {
				carts[i].die, carts[posMap[pos]-1].die, posMap[pos] = true, true, 0
				continue
			}
			posMap[pos] = i + 1
		}
	}
}

func secondChallenge() {
	l := getInput()
	t := genTracks(l)
	adj := genAdjList(t, l)

	carts := getCarts(l)
	simulate2(adj, carts)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
