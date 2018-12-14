package main

import (
	"fmt"
	"os"
)

func getInput() int {
	n := 0
	fmt.Scan(&n)
	return n
}

func firstChallenge() {
	n := getInput()
	r, e1, e2 := []int{3, 7}, 0, 1
	for len(r) < (n + 10) {
		sum := r[e1] + r[e2]
		newR := []int{sum}
		if sum > 9 {
			newR = []int{sum / 10, sum % 10}
		}
		r = append(r, newR...)
		e1, e2 = (r[e1]+e1+1)%len(r), (r[e2]+e2+1)%len(r)
	}
	for _, v := range r[n : n+10] {
		fmt.Print(v)
	}
	fmt.Println()
}

func secondChallenge() {
	s := ""
	fmt.Scan(&s)
	nSlice := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		nSlice[i] = int(s[i] - "0"[0])
	}
	r, e1, e2, pointer := []int{3, 7}, 0, 1, 0
	for {
		sum := r[e1] + r[e2]
		newR := []int{sum}
		if sum > 9 {
			newR = []int{sum / 10, sum % 10}
		}
		r = append(r, newR...)
		e1, e2 = (r[e1]+e1+1)%len(r), (r[e2]+e2+1)%len(r)

		// Can start compare
		if len(r) >= len(nSlice) {
			match := true
			for i := 0; i < len(s); i++ {
				if nSlice[i] != r[pointer+i] {
					match = false
				}
			}
			if match {
				fmt.Println(pointer)
				os.Exit(0)
			}
			pointer++
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
