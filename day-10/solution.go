package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getInput() [][]int {
	res := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a, b, c, d := 0, 0, 0, 0
		fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%d, %d>", &a, &b, &c, &d)
		res = append(res, []int{a, b, c, d})
	}
	return res
}

func getVertiDistance(src [][]int) int {
	maxY, minY := math.MinInt32, math.MaxInt32
	for _, v := range src {
		if v[1] > maxY {
			maxY = v[1]
		}
		if v[1] < minY {
			minY = v[1]
		}
	}
	return maxY - minY + 1
}

func updatePosition(src [][]int) {
	for _, v := range src {
		v[0] += v[2]
		v[1] += v[3]
	}
}

func print(src [][]int) {
	maxY, minY, maxX, minX := math.MinInt32, math.MaxInt32, math.MinInt32, math.MaxInt32
	pointMap := map[[2]int]bool{}
	for _, v := range src {
		if v[0] > maxX {
			maxX = v[0]
		}
		if v[0] < minX {
			minX = v[0]
		}
		if v[1] > maxY {
			maxY = v[1]
		}
		if v[1] < minY {
			minY = v[1]
		}
		pointMap[[2]int{v[0], v[1]}] = true
	}
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if pointMap[[2]int{j, i}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	points := getInput()
	counter := 0
	for {
		vertDistance := getVertiDistance(points)
		if vertDistance < 80 {
			fmt.Println("sep")
			print(points)
			fmt.Println(counter)
		}
		updatePosition(points)
		counter++
	}
}
