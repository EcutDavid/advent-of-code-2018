// Clac the board by building a (MaxX + 1) * (MaxY + 1) grid, the answer should be
// The one who does not touch border but has most points.
// A much simpler solution is simply calc the distances to all cord for all the points on the board...
// Came up this thought after spending too much time on BFS.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getInput() [][2]int {
	scanner := bufio.NewScanner(os.Stdin)
	xyList := [][2]int{}
	for scanner.Scan() {
		x, y := 0, 0
		fmt.Sscanf(scanner.Text(), "%d, %d", &x, &y)
		xyList = append(xyList, [2]int{x, y})
	}
	return xyList
}

var dirs = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}

func bfs(board [][]int, index, width, height, key int) {
	queue, visited := [][2]int{[2]int{index % width, index / width}}, map[int]bool{index: true}
	distance := map[int]int{index: 0}

	for len(queue) > 0 {
		task := queue[0]
		boardIndex := task[1]*width + task[0]
		queue = queue[1:]
		// Collide
		if board[boardIndex][0] > 0 && board[boardIndex][0] == distance[boardIndex] {
			board[boardIndex][1] = -1
		}
		if board[boardIndex][0] > distance[boardIndex] {
			board[boardIndex][0], board[boardIndex][1] = distance[boardIndex], key
		}

		for i := 0; i < len(dirs); i++ {
			newX, newY := task[0]+dirs[i][0], task[1]+dirs[i][1]
			if newX < 0 || newX >= width {
				continue
			}
			if newY < 0 || newY >= height {
				continue
			}
			newBoardIndex := newY*width + newX
			if visited[newBoardIndex] {
				continue
			}
			visited[newBoardIndex] = true
			queue, distance[newBoardIndex] = append(queue, [2]int{newX, newY}), distance[boardIndex]+1
		}
	}
}

func firstChallenge() {
	xyList := getInput()
	maxX, maxY := 0, 0
	for i := 0; i < len(xyList); i++ {
		if xyList[i][0] > maxX {
			maxX = xyList[i][0]
		}
		if xyList[i][1] > maxY {
			maxY = xyList[i][1]
		}
	}
	width := maxX + 1
	board := make([][]int, width*(maxY+1))

	for i := 0; i < len(board); i++ {
		board[i] = []int{math.MaxInt32, -1}
	}
	for i := 0; i < len(xyList); i++ {
		index := xyList[i][0] + xyList[i][1]*width
		bfs(board, index, width, maxY+1, i+1)
	}

	count := map[int]int{}
	badGuy := map[int]bool{}
	for i := 0; i < len(board); i++ {
		x, y := i%width, i/width
		if board[i][1] == -1 {
			continue
		}
		if x == 0 || x == maxX || y == 0 || y == maxY {
			badGuy[board[i][1]] = true
			continue
		}
		count[board[i][1]]++
	}

	max := 0
	for i := 1; i <= len(xyList); i++ {
		if badGuy[i] {
			continue
		}
		if count[i] > max {
			max = count[i]
		}
	}
	fmt.Println(max)
}

func distance(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func secondChallenge() {
	xyList := getInput()
	maxX, maxY := 0, 0
	for i := 0; i < len(xyList); i++ {
		if xyList[i][0] > maxX {
			maxX = xyList[i][0]
		}
		if xyList[i][1] > maxY {
			maxY = xyList[i][1]
		}
	}
	width := maxX + 1

	total, boardLen := 0, width*(maxY+1)
	for i := 0; i < boardLen; i++ {
		x, y := i%width, i/width
		sum := 0
		for i := 0; i < len(xyList); i++ {
			sum += distance(xyList[i], [2]int{x, y})
		}
		if sum < 10000 {
			total++
		}
	}
	fmt.Println(total)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
