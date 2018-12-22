package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	sand = iota
	wall
	spring
	solid
)

// Only for debug
func drawBoard(board [][]int) {
	for _, line := range board {
		for _, v := range line {
			if v == wall {
				fmt.Print("#")
			} else if v == sand {
				fmt.Print(".")
			} else if v == solid {
				fmt.Print("S")
			} else {
				fmt.Print("+")
			}

		}
		fmt.Println()
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func genBoard(xList, yList [][3]int) ([][]int, int) {
	maxX, minX, maxY, minY := 500, 500, 1, 500
	for _, v := range xList {
		maxX, minX = max(v[0], maxX), min(v[0], minX)
		maxY, minY = max(v[2], maxY), min(v[1], minY)
	}
	for _, v := range yList {
		maxX, minX = max(v[2], maxX), min(v[1], minX)
		maxY, minY = max(v[0], maxY), min(v[0], minY)
	}
	maxX, minX = maxX+25, minX-25
	width, height := maxX-minX+1, maxY-0+1

	board := make([][]int, height)
	for i := 0; i < height; i++ {
		board[i] = make([]int, width)
	}

	for _, v := range xList {
		for i := v[1]; i <= v[2]; i++ {
			board[i][v[0]-minX] = wall
		}
	}
	for _, v := range yList {
		for i := v[1]; i <= v[2]; i++ {
			board[v[0]][i-minX] = wall
		}
	}
	board[0][500-minX] = spring
	return board, minY
}

func simulate(board [][]int) {
	h, w := len(board), len(board[0])
	simulatePoint := func(i, j int) {
		if board[i][j] != spring {
			return
		}
		// Generate new spring vertically
		if board[i+1][j] == sand {
			y := i + 1
			for y < h && board[y][j] == sand {
				board[y][j], y = spring, y+1
			}
			return
		}

		// Generate new solid
		veryLeft, veryRight := j, j
		for k := j - 1; k >= 0; k-- {
			if board[i][k] == wall {
				break
			}
			veryLeft = k
		}
		for k := j + 1; k < w; k++ {
			if board[i][k] == wall {
				break
			}
			veryRight = k
		}
		canFill := true
		for k := veryLeft; k <= veryRight; k++ {
			// A new solid can only grow from wall or other solid
			if board[i+1][k] != solid && board[i+1][k] != wall {
				canFill = false
			}
		}
		if canFill {
			for k := veryLeft; k <= veryRight; k++ {
				board[i][k] = solid
			}
			return
		}

		// Generate new spring that "overflow"
		veryLeft, veryRight = j, j
		for k := j - 1; k >= 0; k-- {
			if board[i][k] == wall {
				break
			}
			if board[i+1][k] != sand {
				veryLeft = k
			} else if board[i+1][k+1] == wall {
				veryLeft = k
				break
			} else {
				break
			}
		}
		for k := j + 1; k < w; k++ {
			if board[i][k] == wall {
				break
			}
			if board[i+1][k] != sand {
				veryRight = k
			} else if board[i+1][k-1] == wall {
				veryRight = k
				break
			} else {
				break
			}
		}
		if veryLeft == veryRight {
			return
		}

		canFill = true
		for k := veryLeft; k <= veryRight; k++ {
			if k == veryLeft && board[i+1][k+1] == wall {
				continue
			}
			if k == veryRight && board[i+1][k-1] == wall {
				continue
			}
			// A new spring can only grow from wall or solid, except the very left and very right one
			if board[i+1][k] != solid && board[i+1][k] != wall {
				canFill = false
			}
		}
		if !canFill {
			return
		}
		for k := veryLeft; k <= veryRight; k++ {
			board[i][k] = spring
		}
	}

	//Just a random number big enough,
	//each round(f) Fu
	for r := 0; r < 9000; r++ {
		for i := 0; i < (h - 1); i++ {
			for j := 0; j < w; j++ {
				simulatePoint(i, j)
			}
		}
	}
}

func firstChallenge(xList, yList [][3]int) {
	board, minY := genBoard(xList, yList)
	simulate(board)
	h, w, sum := len(board), len(board[0]), 0
	for i := minY; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == spring || board[i][j] == solid {
				sum++
			}
		}
	}
	// drawBoard(board)
	fmt.Println(sum)
}

func secondChallenge(xList, yList [][3]int) {
	board, _ := genBoard(xList, yList)
	simulate(board)
	h, w, sum := len(board), len(board[0]), 0
	for i := 1; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == solid {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func parseInput() ([][3]int, [][3]int) {
	xList, yList := [][3]int{}, [][3]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		a, b, c, d, e := "", "", 0, 0, 0
		fmt.Sscanf(line, "%1s=%d, %1s=%d..%d", &a, &c, &b, &d, &e)
		if a == "x" {
			xList = append(xList, [3]int{c, d, e})
		} else {
			yList = append(yList, [3]int{c, d, e})
		}
	}
	return xList, yList
}

func main() {
	l1, l2 := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(l1, l2)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(l1, l2)
}
