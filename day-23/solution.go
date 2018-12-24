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

func secondChallenge(bots []robot) {
	bestSum, bestDistance, targetBots, bestVertex, bestBotIndex := 0, int64(math.MaxInt64), []int{}, [3]int64{}, 0
	for k, b := range bots {
		for _, v := range genVertices(b) {
			sum, distanceToV, potentialBots := 0, getDistance(origin, v), []int{}
			for k, d := range bots {
				if getDistance(d.pos, v) <= d.r {
					sum++
					potentialBots = append(potentialBots, k)
				}
			}
			if sum > bestSum {
				fmt.Println(v)
				bestDistance, bestSum, targetBots, bestVertex = distanceToV, sum, potentialBots, v
				fmt.Println(b)
				bestBotIndex = k
			} else if (sum == bestSum) && (bestDistance > distanceToV) {
				bestDistance = distanceToV
				fmt.Println("concern", v)
			}
		}
	}

	slack := int64(math.MaxInt64)
	for k, i := range targetBots {
		if k == bestBotIndex {
			continue
		}
		slack = min(slack, bots[i].r-getDistance(bots[i].pos, bestVertex))
	}
	fmt.Println(bestVertex, bots[bestBotIndex], slack, getDistance(origin, bestVertex))

	// The question is, can this best Vertext be moved somewhere eles?
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
