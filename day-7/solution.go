package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getInput() [][2]string {
	scanner, s := bufio.NewScanner(os.Stdin), [][2]string{}
	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), " ")
		a, b := raw[1], raw[len(raw)-3]
		s = append(s, [2]string{a, b})
	}
	return s
}

func buildAdjList(s [][2]string) map[string]map[string]bool {
	list := map[string]map[string]bool{}
	for _, v := range s {
		if list[v[0]] == nil {
			list[v[0]] = map[string]bool{}
		}
		list[v[0]][v[1]] = true
	}
	return list
}

func topSort(s [][2]string, adjList map[string]map[string]bool) []string {
	inDegree, p := map[string]int{}, map[string]bool{}
	for _, v := range s {
		inDegree[v[1]], p[v[0]], p[v[1]] = inDegree[v[1]]+1, true, true
	}

	queue := []string{}
	for k := range p {
		if inDegree[k] == 0 {
			queue = append(queue, k)
		}
	}

	path := []string{}
	for len(queue) > 0 {
		sort.Strings(queue)
		task := queue[0]
		queue, path = queue[1:], append(path, task)
		for k := range adjList[task] {
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	return path
}

// Assume s is 1 char long
func getDistance(s string, fir int) int {
	return int(s[0]-"A"[0]+1) + fir
}

type queueItem struct {
	tar       string
	distance  int
	hasWorker bool
}
type taskQueue []queueItem

func (t taskQueue) Len() int {
	return len(t)
}
func (t taskQueue) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t taskQueue) Less(i, j int) bool {
	if t[i].distance != t[j].distance {
		return t[i].distance < t[j].distance
	}
	return t[i].tar < t[j].tar
}

func calcDistance(s [][2]string, adjList map[string]map[string]bool) int {
	inDegree, p, factor, workerCount := map[string]int{}, map[string]bool{}, 60, 5
	for _, v := range s {
		inDegree[v[1]], p[v[0]], p[v[1]] = inDegree[v[1]]+1, true, true
	}

	queue := taskQueue{}
	for k := range p {
		if inDegree[k] == 0 {
			queue = append(queue, queueItem{k, getDistance(k, factor), false})
		}
	}
	for i := 0; i < len(queue); i++ {
		if i >= workerCount {
			break
		}
		queue[i].hasWorker = true
	}

	totalTime := 0
	for len(queue) > 0 {
		sort.Sort(queue)
		task := queue[0]
		for i := 1; i < len(queue); i++ {
			if queue[i].hasWorker {
				queue[i].distance -= task.distance
			}
		}

		queue, totalTime = queue[1:], totalTime+task.distance

		for k := range adjList[task.tar] {
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = append(queue, queueItem{k, getDistance(k, factor), false})
			}
		}

		extraWorker, itemsWithoutWorker, tarIndexMap := workerCount, []string{}, map[string]int{}
		for i := 0; i < len(queue); i++ {
			if queue[i].hasWorker {
				extraWorker--
			} else {
				itemsWithoutWorker, tarIndexMap[queue[i].tar] = append(itemsWithoutWorker, queue[i].tar), i
			}
		}
		sort.Strings(itemsWithoutWorker)
		for i := 0; i < (len(itemsWithoutWorker)) && (extraWorker > 0); i++ {
			queue[tarIndexMap[itemsWithoutWorker[i]]].hasWorker = true
			extraWorker--
		}
	}
	return totalTime
}

func firstChallenge() {
	src := getInput()
	adj := buildAdjList(src)
	path := topSort(src, adj)
	fmt.Println(strings.Join(path, ""))
}

func secondChallenge() {
	src := getInput()
	adj := buildAdjList(src)
	time := calcDistance(src, adj)
	fmt.Println(time)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
