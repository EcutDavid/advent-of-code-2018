package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	sand = iota
	wall
	water
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
	maxX, minX = maxX+1, minX-1
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
	board[0][500-minX] = water
	return board
}

func simulate(board [][]int) {
	h, w := len(board), len(board[0])
	for {
		// As long as there are water in the last layer, we end
		end := false
		for i := 0; i < w; i++ {
			if board[h-1][i] == water {
				end = true
			}
		}
		if end {
			break
		}

		hasFlowY := false
		for i := 0; i < (h - 1); i++ {
			for j := 0; j < w; j++ {
				if board[i][j] != water {
					continue
				}
				if board[i+1][j] == sand {
					hasFlowY = true
					y := i + 1
					for y < h && board[y][j] == sand {
						board[y][j], y = water, y+1
					}
					continue
				}
			}
		}
		// Cannot grow horizontally
		if hasFlowY {
			// drawBoard(board)
			continue
		}
		filledSolid := false
		for i := (h - 1); i >= 0; i-- {
			for j := 0; j < w; j++ {
				if board[i][j] != water {
					continue
				}
				veryLeft, veryRight := j, j
				for k := j - 1; k >= 0; k-- {
					if board[i][k] == wall {
						break
					}
					if board[i+1][k] != sand || board[i+1][k+1] == wall {
						veryLeft--
					} else {
						break
					}
				}
				for k := j + 1; k < w; k++ {
					if board[i][k] == wall {
						break
					}
					if board[i+1][k] != sand || board[i+1][k-1] == wall {
						veryRight++
					} else {
						break
					}
				}
				// fmt.Println(i, veryLeft, veryRight)
				canFill := true
				for k := veryLeft; k <= veryRight; k++ {
					if board[i+1][k] == sand {
						canFill = false
					}
				}
				if canFill {
					for k := veryLeft; k <= veryRight; k++ {
						if board[i][k] != water {
							filledSolid = true
						}
						board[i][k] = solid
					}
				}
			}
		}

		// Need more water
		if filledSolid {
			continue
		}
		filledWater := false
		for i := (h - 1); i >= 0; i-- {
			for j := 0; j < w; j++ {
				if board[i][j] != water {
					continue
				}
				veryLeft, veryRight := j, j
				for k := j - 1; k >= 0; k-- {
					if board[i][k] == wall {
						break
					}
					if board[i+1][k] != sand || board[i+1][k+1] == wall {
						veryLeft--
					} else {
						break
					}
				}
				for k := j + 1; k < w; k++ {
					if board[i][k] == wall {
						break
					}
					if board[i+1][k] != sand || board[i+1][k-1] == wall {
						veryRight++
					} else {
						break
					}
				}
				// fmt.Println(i, veryLeft, veryRight)
				for k := veryLeft; k <= veryRight; k++ {
					board[i][k] = water
					if board[i][k] != water {
						filledWater = true
					}
				}
			}
			if filledWater {
				break
			}
		}
		// drawBoard(board)
	}
}

func firstChallenge(xList, yList [][3]int) {
	board := genBoard(xList, yList)
	simulate(board)
	h, w, sum := len(board), len(board[0]), 0
	for i := 1; i < h; i++ {
		for j := 1; j < (w - 1); j++ {
			if board[i][j] == water {
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
