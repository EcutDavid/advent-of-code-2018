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

// Debug only
func drawBoard(board [][]int) {
	fmt.Println("A")
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
	fmt.Println("B")
}

func parseInput() ([][3]int, [][3]int) {
	xList, yList := [][3]int{}, [][3]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		a, b, c, d, e := "", "", 0, 0, 0
		fmt.Sscanf(line, "%1s=%d, %1s=%d..%d", &a, &c, &b, &d, &e)
		if a == "x" {
			// if e > 100 {
			// 	continue
			// }
			xList = append(xList, [3]int{c, d, e})
		} else {
			// if c > 100 {
			// 	continue
			// }
			yList = append(yList, [3]int{c, d, e})
		}
	}
	return xList, yList
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

func genBoard(xList, yList [][3]int) [][]int {
	maxX, minX, maxY, minY := 500, 500, 1, 0
	for _, v := range xList {
		maxX, minX = max(v[0], maxX), min(v[0], minX)
		maxY = max(v[2], maxY)
	}
	for _, v := range yList {
		maxX, minX = max(v[2], maxX), min(v[1], minX)
		maxY = max(v[0], maxY)
	}
	maxX, minX = maxX+25, minX-25
	width, height := maxX-minX+1, maxY-minY+1

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
	return board
}

func simulate(board [][]int) {
	h, w := len(board), len(board[0])
	//TODO: improve max
	for f := 0; f < 9000; f++ {
		for i := 0; i < (h - 1); i++ {
			for j := 0; j < w; j++ {
				if board[i][j] != spring {
					continue
				}
				if board[i+1][j] == sand {
					y := i + 1
					for y < h && board[y][j] == sand {
						board[y][j], y = spring, y+1
					}
					continue
				}
			}
		}

		for i := (h - 2); i >= 0; i-- {
			for j := 0; j < w; j++ {
				if board[i][j] != spring {
					continue
				}
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
				if !canFill {
					continue
				}
				for k := veryLeft; k <= veryRight; k++ {
					board[i][k] = solid
				}
			}
		}

		for i := (h - 2); i >= 0; i-- {
			for j := 0; j < w; j++ {
				if board[i][j] != spring {
					continue
				}
				veryLeft, veryRight := j, j
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
					continue
				}

				canFill := true
				for k := veryLeft; k <= veryRight; k++ {
					// fmt.Println(w, h, i+1, k+1, veryLeft, veryRight)
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
					continue
				}
				for k := veryLeft; k <= veryRight; k++ {
					board[i][k] = spring
				}
			}
		}
	}
}

func firstChallenge(xList, yList [][3]int) {
	board := genBoard(xList, yList)
	simulate(board)
	h, w, sum := len(board), len(board[0]), 0
	for i := 1; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == spring || board[i][j] == solid {
				sum++
			}
		}
	}
	drawBoard(board)
	fmt.Println(sum)
}

func secondChallenge() {

}

func main() {
	l1, l2 := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(l1, l2)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge()
}
