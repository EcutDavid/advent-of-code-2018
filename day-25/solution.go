package main

import (
	"bufio"
	"fmt"
	"os"
)

type spaceTime [4]int

func (s1 *spaceTime) isAdj(s2 *spaceTime) bool {
	return getDistance(s1, s2) <= 3
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getDistance(a, b *spaceTime) int {
	var sum int
	for i := 0; i < 4; i++ {
		sum += abs(a[i] - b[i])
	}
	return sum
}

func challenge(list []*spaceTime) {
	adjMap, visited, groupSum := map[int]map[int]bool{}, map[int]bool{}, 0
	for i := 0; i < len(list); i++ {
		adjMap[i] = map[int]bool{}
		for j := 0; j < len(list); j++ {
			if i == j {
				continue
			}
			if list[i].isAdj(list[j]) {
				adjMap[i][j] = true
			}
		}
	}

	search := func(root int) {
		queue := []int{root}
		for len(queue) > 0 {
			task := queue[0]
			queue = queue[1:]

			for d := range adjMap[task] {
				if visited[d] {
					continue
				}
				queue, visited[d] = append(queue, d), true
			}
		}
	}

	for i := 0; i < len(list); i++ {
		if visited[i] {
			continue
		}
		visited[i], groupSum = true, groupSum+1
		search(i)
	}
	fmt.Println(groupSum)
}

func parseInput() []*spaceTime {
	list := []*spaceTime{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, z, t int
		fmt.Sscanf(line, "%d,%d,%d,%d", &x, &y, &z, &t)
		list = append(list, &spaceTime{x, y, z, t})
	}
	return list
}

func main() {
	list := parseInput()
	challenge(list)
}
