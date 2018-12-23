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

func genVertices(bot robot) [7][3]int64 {
	return [7][3]int64{
		[3]int64{bot.x + bot.r, bot.y, bot.z},
		[3]int64{bot.x - bot.r, bot.y, bot.z},
		[3]int64{bot.x, bot.y + bot.r, bot.z},
		[3]int64{bot.x, bot.y - bot.r, bot.z},
		[3]int64{bot.x, bot.y, bot.z + bot.r},
		[3]int64{bot.x, bot.y, bot.z - bot.r},
		[3]int64{bot.x, bot.y, bot.z},
	}
}
func getDistance(r robot, pos [3]int64) int64 {
	dx, dy, dz := r.x-pos[0], r.y-pos[1], r.z-pos[2]
	return abs(dx) + abs(dy) + abs(dz)
}

func secondChallenge(bots []robot) {
	bestSum, bestDistance := 0, int64(math.MaxInt64)
	for _, b := range bots {
		for _, v := range genVertices(b) {
			sum, distanceToV := 0, getDistance(origin, v)
			for _, d := range bots {
				if getDistance(d, v) <= d.r {
					sum++
				}
			}
			if sum > bestSum {
				fmt.Println(v)
				bestDistance, bestSum = distanceToV, sum
			} else if (sum == bestSum) && (bestDistance > distanceToV) {
				bestDistance = distanceToV
				fmt.Println(v)
			}
		}
	}
	fmt.Println(bestSum, bestDistance)
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
