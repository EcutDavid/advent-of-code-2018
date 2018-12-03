package main

import (
	"bufio"
	"fmt"
	"os"
)

func firstChallenge() {
	scanner := bufio.NewScanner(os.Stdin)
	recs := [][4]int{}

	for scanner.Scan() {
		n, x, y, w, h := 0, 0, 0, 0, 0
		line := scanner.Text()
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &n, &x, &y, &w, &h)
		recs = append(recs, [4]int{x, y, w, h})
	}

	maxX, maxY := 0, 0
	for _, v := range recs {
		x, y := v[0]+v[2], v[1]+v[3]
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
	}
	whole := make([]int, (maxX+1)*(maxY+1))
	for _, v := range recs {
		for i := 0; i < v[2]; i++ {
			for j := 0; j < v[3]; j++ {
				index := (j+v[1])*maxX + i + v[0]
				// fmt.Println(index, (j + v[1]), i+v[0])
				whole[index]++
			}
		}
	}
	for i := 0; i < maxY; i++ {
		// fmt.Println(whole[i*maxX : (i+1)*maxX])
	}

	sum := 0
	for _, v := range whole {
		if v > 1 {
			sum++
		}
	}
	fmt.Println(sum)
}

func secondChallenge() {
	scanner := bufio.NewScanner(os.Stdin)
	recs := [][4]int{}

	for scanner.Scan() {
		n, x, y, w, h := 0, 0, 0, 0, 0
		line := scanner.Text()
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &n, &x, &y, &w, &h)
		recs = append(recs, [4]int{x, y, w, h})
	}

	maxX, maxY := 0, 0
	for _, v := range recs {
		x, y := v[0]+v[2], v[1]+v[3]
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
	}
	whole := make([]int, (maxX+1)*(maxY+1))
	for _, v := range recs {
		for i := 0; i < v[2]; i++ {
			for j := 0; j < v[3]; j++ {
				index := (j+v[1])*maxX + i + v[0]
				// fmt.Println(index, (j + v[1]), i+v[0])
				whole[index]++
			}
		}
	}

	for n, v := range recs {
		theOne := true
		for i := 0; i < v[2]; i++ {
			for j := 0; j < v[3]; j++ {
				index := (j+v[1])*maxX + i + v[0]
				if whole[index] > 1 {
					theOne = false
				}
			}
		}
		if theOne {
			fmt.Println(n + 1)
			os.Exit(0)
		}
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
