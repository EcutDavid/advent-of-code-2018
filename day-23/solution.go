// Still work in progress, current solution has a hole in logic which I still didn't find out where
// Sorry for bad code, will clean up after finish the solution(if).
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type robot struct {
	x, y, z int64
	r       int64
}

var origin = robot{0, 0, 0, 0}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func getBotDistance(a, b robot) int64 {
	dx, dy, dz := a.x-b.x, a.y-b.y, a.z-b.z
	return abs(dx) + abs(dy) + abs(dz)
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
		if getBotDistance(bots[i], bots[bestR]) <= bots[bestR].r {
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

func genVertices(bot robot) [6][3]int64 {
	return [6][3]int64{
		[3]int64{bot.x + bot.r, bot.y, bot.z},
		[3]int64{bot.x - bot.r, bot.y, bot.z},
		[3]int64{bot.x, bot.y + bot.r, bot.z},
		[3]int64{bot.x, bot.y - bot.r, bot.z},
		[3]int64{bot.x, bot.y, bot.z + bot.r},
		[3]int64{bot.x, bot.y, bot.z - bot.r},
	}
}
func getDistance(r robot, pos [3]int64) int64 {
	dx, dy, dz := r.x-pos[0], r.y-pos[1], r.z-pos[2]
	return abs(dx) + abs(dy) + abs(dz)
}

// func genBotRange(bot robot) [3][2]int64 {
// 	return [3][2]int64{
// 		[2]int64{bot.x - bot.r, bot.x + bot.r},
// 		[2]int64{bot.y - bot.r, bot.y + bot.r},
// 		[2]int64{bot.z - bot.r, bot.y + bot.r},
// 	}
// }
// func genMergedRange(a, b [3][2]int64) [3][2]int64 {
// 	return [3][2]int64{
// 		[2]int64{max(a[0][0], b[0][0]), min(a[0][1], b[0][1])},
// 		[2]int64{max(a[1][0], b[1][0]), min(a[1][1], b[1][1])},
// 		[2]int64{max(a[2][0], b[2][0]), min(a[2][1], b[2][1])},
// 	}
// }

func secondChallenge(bots []robot) {
	bestSum, bestDistance, targetBots, bestVertex, bestBotIndex := 0, int64(math.MaxInt64), []int{}, [3]int64{}, 0
	for k, b := range bots {
		for _, v := range genVertices(b) {
			sum, distanceToV, potentialBots := 0, getDistance(origin, v), []int{}
			for k, d := range bots {
				if getDistance(d, v) <= d.r {
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
		slack = min(slack, bots[i].r-getDistance(bots[i], bestVertex))
	}
	fmt.Println(bestVertex, bots[bestBotIndex], slack, getDistance(origin, bestVertex))

	// The question is, can this best Vertext be moved somewhere eles?
}

// func secondChallenge(bots []robot) {
// 	groups := [][]robot{}
// 	// Generate groups that in each group, all bots can "reach" each other
// 	for i := 0; i < len(bots); i++ {
// 		groups = append(groups, []robot{})
// 		for j := 0; j < len(bots); j++ {
// 			if getDistance(bots[i], bots[j]) <= (bots[i].r + bots[j].r) {
// 				groups[i] = append(groups[i], bots[j])
// 			}
// 		}
// 	}

// 	// Filter groups that don't have enough bots
// 	bestGroupLength := 0
// 	for _, g := range groups {
// 		bestGroupLength = max(len(g), bestGroupLength)
// 	}
// 	choosedGroups := [][]robot{}
// 	for _, g := range groups {
// 		if len(g) == bestGroupLength {
// 			choosedGroups = append(choosedGroups, g)
// 		}
// 	}
// 	best := math.MaxInt32
// 	for _, v := range choosedGroups {
// 		fmt.Println(v)
// 		// Find the bot that cloest to origin, and check how much can it be more closely to origin.
// 		theBotIndex, shortest := 0, math.MaxInt32
// 		for i := 0; i < len(v); i++ {
// 			distance := getDistance(origin, v[i])
// 			if shortest > distance {
// 				theBotIndex, shortest = i, distance
// 			}
// 		}
// 		theBot := v[theBotIndex]
// 		minR := -theBot.r
// 		for _, b := range v {
// 			distance := getDistance(b, theBot)
// 			// have to increase minR to meet the requirement.
// 			if distance > (b.r + minR) {
// 				minR = distance - b.r
// 			}
// 		}
// 		if best > (shortest + minR) {
// 			best = shortest + minR
// 		}
// 	}
// 	fmt.Println(best)
// }

func parseInput() []robot {
	bots := []robot{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, z, r int64
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		bots = append(bots, robot{x, y, z, r})
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
