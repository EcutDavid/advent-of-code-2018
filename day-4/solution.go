package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type lineMeta = struct {
	guard int
	time  int
}

func parseInput() []lineMeta {
	scanner := bufio.NewScanner(os.Stdin)
	rawLines := []string{}
	for scanner.Scan() {
		rawLines = append(rawLines, scanner.Text())
	}
	sort.Strings(rawLines)
	lines := []lineMeta{}
	for _, line := range rawLines {
		var t1, t2, t3, t4, t5 int64 = 0, 0, 0, 0, 0
		fmt.Sscanf(line[:19], "[%d-%d-%d %d:%d]", &t1, &t2, &t3, &t4, &t5)

		if line[19:19+5] == "Guard" {
			words := strings.Split(line[19:], " ")
			id, _ := strconv.Atoi(words[1][1:])
			lines = append(lines, lineMeta{
				guard: id,
			})
		} else {
			lines = append(lines, lineMeta{
				time:  int(t5),
				guard: -1,
			})
		}
	}
	return lines
}

func firstChallenge(src []lineMeta) {
	guards, curGuard := map[int][][2]int{}, -1
	for i := 0; i < len(src); i++ {
		if src[i].guard > -1 {
			curGuard = src[i].guard
			if guards[curGuard] == nil {
				guards[curGuard] = [][2]int{}
			}
		} else {
			guards[curGuard], i = append(guards[curGuard], [2]int{src[i].time, src[i+1].time}), i+1
		}
	}

	sleepMostID, most := -1, 0
	for id, v := range guards {
		time := 0
		for _, T := range v {
			time += T[1] - T[0] + 1
		}
		if time > most {
			most, sleepMostID = time, id
		}
	}

	theGuard := guards[sleepMostID]
	minMap := map[int]int{}
	for _, T := range theGuard {
		for i := T[0]; i < T[1]; i++ {
			minMap[i]++
		}
	}
	best := -1
	for t, c := range minMap {
		if c > minMap[best] {
			best = t
		}
	}
	fmt.Println(best * sleepMostID)
}

func secondChallenge(src []lineMeta) {
	guards := map[int][][2]int{}
	curGuard := -1
	for i := 0; i < len(src); i++ {
		if src[i].guard > -1 {
			curGuard = src[i].guard
			if guards[curGuard] == nil {
				guards[curGuard] = [][2]int{}
			}
		} else {
			guards[curGuard], i = append(guards[curGuard], [2]int{src[i].time, src[i+1].time}), i+1
		}
	}
	minCountSlice := [][3]int{}

	for id, g := range guards {
		minMap := map[int]int{}
		for _, T := range g {
			for i := T[0]; i < T[1]; i++ {
				minMap[i]++
			}
		}
		best := -1
		for t, c := range minMap {
			if c > minMap[best] {
				best = t
			}
		}
		minCountSlice = append(minCountSlice, [3]int{id, best, minMap[best]})
	}
	best := 0
	for t, c := range minCountSlice {
		if c[2] > minCountSlice[best][2] {
			best = t
		}
	}
	fmt.Println(minCountSlice[best][0] * minCountSlice[best][1])

}

func main() {
	input := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(input)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(input)
}
