// Still work in progress, current solution has a hole in logic which I still didn't find out where
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type robot struct {
	pos [3]int64
	r   int64
}

var origin = [3]int64{0, 0, 0}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func genVertices(bot robot) [][3]int64 {
	return [][3]int64{
		[3]int64{bot.pos[0] + bot.r, bot.pos[1], bot.pos[2]},
		[3]int64{bot.pos[0] - bot.r, bot.pos[1], bot.pos[2]},
		[3]int64{bot.pos[0], bot.pos[1] + bot.r, bot.pos[2]},
		[3]int64{bot.pos[0], bot.pos[1] - bot.r, bot.pos[2]},
		[3]int64{bot.pos[0], bot.pos[1], bot.pos[2] + bot.r},
		[3]int64{bot.pos[0], bot.pos[1], bot.pos[2] - bot.r},
	}
}
func getDistance(a, b [3]int64) int64 {
	var sum int64
	for i := 0; i < 3; i++ {
		sum += abs(a[i] - b[i])
	}
	return sum
}

func firstChallenge(bots []robot) {
	bestR, best := -1, int64(0)
	for i := 0; i < len(bots); i++ {
		if bots[i].r > best {
			best, bestR = bots[i].r, i
		}
	}
	sum := 0
	for i := 0; i < len(bots); i++ {
		if getDistance(bots[i].pos, bots[bestR].pos) <= bots[bestR].r {
			sum++
		}
	}
	fmt.Println(sum)
}

func secondChallenge(bots []robot) {
	bestSum, targetBots, bestVertex := 0, []int{}, [3]int64{}
	for _, b := range bots {
		for _, v := range genVertices(b) {
			sum, potentialBots := 0, []int{}
			for k, d := range bots {
				if getDistance(d.pos, v) <= d.r {
					sum++
					potentialBots = append(potentialBots, k)
				}
			}
			if sum > bestSum {
				bestSum, targetBots, bestVertex = sum, potentialBots, v
			}
		}
	}
	vertexMap := map[[3]int64]bool{}
	for _, b := range bots {
		for _, v := range genVertices(b) {
			sum := 0
			for _, d := range bots {
				if getDistance(d.pos, v) <= d.r {
					sum++
				}
			}
			if sum == bestSum {
				if vertexMap[v] {
					continue
				}
				vertexMap[v] = true
				fmt.Println("potential vertex", v)
			}
		}
	}

	validate := func(testVertex [3]int64) bool {
		for _, index := range targetBots {
			if getDistance(testVertex, bots[index].pos) > bots[index].r {
				return false
			}
		}
		return true
	}

	queue, visited := [][3]int64{bestVertex}, map[[3]int64]bool{bestVertex: true}
	dirs, bestDistance := [][3]int64{[3]int64{0, 0, -1}, [3]int64{-1, 0, 0}, [3]int64{0, -1, 0}}, int64(math.MaxInt64)
	for len(queue) > 0 {
		task, distance := queue[0], getDistance(queue[0], origin)
		if distance < bestDistance {
			bestDistance = distance
		}
		queue = queue[1:]
		for _, d := range dirs {
			newPos := [3]int64{task[0] + d[0], task[1] + d[1], task[2] + d[2]}
			if visited[newPos] || !validate(newPos) {
				continue
			}
			visited[newPos], queue = true, append(queue, newPos)
		}
	}

	fmt.Println("final deal", bestDistance)
}

func parseInput() []robot {
	bots := []robot{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, z, r int64
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		bots = append(bots, robot{[3]int64{x, y, z}, r})
	}
	return bots
}

func main() {
	bots := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(bots)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(bots)
}
