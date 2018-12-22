package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	open = iota
	lumberyard
	trees
)

var (
	gASC = "."[0]
	lASC = "#"[0]
	tASC = "|"[0]
)

var dirs = [8][2]int{
	[2]int{0, 1},
	[2]int{0, -1},
	[2]int{1, 1},
	[2]int{1, -1},
	[2]int{-1, 1},
	[2]int{-1, -1},
	[2]int{1, 0},
	[2]int{-1, 0},
}

func getAdj(board [][]int, pos [2]int) [3]int {
	res, w, h := [3]int{}, len(board[0]), len(board)
	for _, dir := range dirs {
		x, y := pos[1]+dir[1], pos[0]+dir[0]
		if (x < 0) || (x >= w) || (y < 0) || (y >= h) {
			continue
		}
		(&res)[board[y][x]]++
	}
	return res
}

func simulation(board [][]int, round int) ([][]int, int, int, map[int][2500]int) {
	w, h := len(board[0]), len(board)
	snapshotPlace, snapShotMap := map[[2500]int]int{}, map[int][2500]int{}

	for r := 0; r < round; r++ {
		snapshot := genSnapshot(board)
		if oldR, ok := snapshotPlace[snapshot]; ok {
			return nil, oldR, r, snapShotMap
		}
		snapshotPlace[snapshot], snapShotMap[r] = r, snapshot

		boardCopy := make([][]int, h)
		for i := 0; i < h; i++ {
			boardCopy[i] = make([]int, w)
			for j := 0; j < w; j++ {
				adj := getAdj(board, [2]int{i, j})
				switch board[i][j] {
				case open:
					if adj[trees] >= 3 {
						boardCopy[i][j] = trees
					} else {
						boardCopy[i][j] = open
					}
				case trees:
					if adj[lumberyard] >= 3 {
						boardCopy[i][j] = lumberyard
					} else {
						boardCopy[i][j] = trees
					}
				case lumberyard:
					if adj[lumberyard] == 0 || adj[trees] == 0 {
						boardCopy[i][j] = open
					} else {
						boardCopy[i][j] = lumberyard
					}
				}
			}
		}

		board = boardCopy
	}
	return board, 0, 0, nil
}

func genSnapshot(board [][]int) [2500]int {
	w, h := len(board[0]), len(board)
	res := [2500]int{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			res[i*w+j] = board[i][j]
		}
	}
	return res
}

func firstChallenge(board [][]int) {
	board, _, _, _ = simulation(board, 10)
	lCounter, tCounter := 0, 0
	for _, r := range board {
		for _, v := range r {
			if v == lumberyard {
				lCounter++
			}
			if v == trees {
				tCounter++
			}
		}
	}
	fmt.Println(lCounter, tCounter, lCounter*tCounter)
}

func secondChallenge(board [][]int) {
	_, l, r, snapshotMap := simulation(board, int(1e9))
	w := r - l
	index := l + (int(1e9)-r)%w
	lCounter, tCounter := 0, 0
	for _, v := range snapshotMap[index] {
		if v == lumberyard {
			lCounter++
		}
		if v == trees {
			tCounter++
		}
	}
	fmt.Println(lCounter, tCounter, lCounter*tCounter)
}

func parseInput() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	board := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case gASC:
				row[i] = open
			case lASC:
				row[i] = lumberyard
			case tASC:
				row[i] = trees
			}
		}
		board = append(board, row)
	}
	return board
}

func main() {
	board := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(board)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(board)
}
