package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type robot struct {
	x, y, z int
	r       int
}

// Just for calc the distance between a bot & origin point.
var origin = robot{0, 0, 0, 0}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getDistance(a, b robot) int {
	dx, dy, dz := a.x-b.x, a.y-b.y, a.z-b.z
	return abs(dx) + abs(dy) + abs(dz)
}

func firstChallenge(bots []robot) {
	bestR, best := -1, 0
	for i := 0; i < len(bots); i++ {
		if bots[i].r > best {
			best, bestR = bots[i].r, i
		}
	}
	sum := 0
	for i := 0; i < len(bots); i++ {
		if getDistance(bots[i], bots[bestR]) <= bots[bestR].r {
			sum++
		}
	}
	fmt.Println(sum)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func secondChallenge(bots []robot) {
	mMin, mMax := math.MaxInt32, 0
	for _, b := range bots {
		dToOrigin := getDistance(b, origin)
		mMin = max(0, min(dToOrigin-b.r, mMin))
		mMax = max(dToOrigin+b.r, mMax)
	}
	fmt.Println(mMin, mMax, mMax-mMin)
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
		x, y, z, r := 0, 0, 0, 0
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
