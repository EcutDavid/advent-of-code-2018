package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// type Interface interface {
//         // Len is the number of elements in the collection.
//         Len() int
//         // Less reports whether the element with
//         // index i should sort before the element with index j.
//         Less(i, j int) bool
//         // Swap swaps the elements with indexes i and j.
//         Swap(i, j int)
// }

type lineMeta = struct {
	rank  int64
	guard int
	time  int
}
type lineMetaList []lineMeta

func (list lineMetaList) Len() int {
	return len(list)
}
func (list lineMetaList) Less(i, j int) bool {
	return list[i].rank < list[j].rank
}
func (list lineMetaList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func firstChallenge() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := lineMetaList{}
	for scanner.Scan() {
		line := scanner.Text()
		var t1, t2, t3, t4, t5 int64 = 0, 0, 0, 0, 0

		fmt.Sscanf(line[:19], "[%d-%d-%d %d:%d]", &t1, &t2, &t3, &t4, &t5)
		rank := t2*3600*1000 + t3*3600 + t4*60 + t5
		if line[19:19+5] == "Guard" {
			words := strings.Split(line[19:], " ")
			id, _ := strconv.Atoi(words[1][1:])
			lines = append(lines, lineMeta{
				rank:  rank,
				guard: id,
			})
		} else {
			lines = append(lines, lineMeta{
				rank:  rank,
				time:  int(t5),
				guard: -1,
			})
		}
	}
	sort.Sort(lines)
	guards := map[int][][2]int{}
	curGuard := -1
	count := 0
	for i := 0; i < len(lines); i++ {
		if lines[i].guard > -1 {
			curGuard = lines[i].guard
			if count%2 == 1 {
			}
			if guards[curGuard] == nil {
				guards[curGuard] = [][2]int{}
			}
			count = 0
		} else {
			guards[curGuard], i = append(guards[curGuard], [2]int{lines[i].time, lines[i+1].time}), i+1
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

func secondChallenge() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := lineMetaList{}
	for scanner.Scan() {
		line := scanner.Text()
		var t1, t2, t3, t4, t5 int64 = 0, 0, 0, 0, 0

		fmt.Sscanf(line[:19], "[%d-%d-%d %d:%d]", &t1, &t2, &t3, &t4, &t5)
		rank := t2*3600*1000 + t3*3600 + t4*60 + t5
		if line[19:19+5] == "Guard" {
			words := strings.Split(line[19:], " ")
			id, _ := strconv.Atoi(words[1][1:])
			lines = append(lines, lineMeta{
				rank:  rank,
				guard: id,
			})
		} else {
			lines = append(lines, lineMeta{
				rank:  rank,
				time:  int(t5),
				guard: -1,
			})
		}
	}
	sort.Sort(lines)
	guards := map[int][][2]int{}
	curGuard := -1
	count := 0
	for i := 0; i < len(lines); i++ {
		if lines[i].guard > -1 {
			curGuard = lines[i].guard
			if count%2 == 1 {
			}
			if guards[curGuard] == nil {
				guards[curGuard] = [][2]int{}
			}
			count = 0
		} else {
			guards[curGuard], i = append(guards[curGuard], [2]int{lines[i].time, lines[i+1].time}), i+1
		}
	}
	// Each guard, produce: [id, time, count]
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
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
