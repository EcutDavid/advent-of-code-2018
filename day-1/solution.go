package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var plus = "+"[0]

func parseInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nums := []int{}
	for scanner.Scan() {
		l := scanner.Text()
		num, _ := strconv.Atoi(l[1:])
		if l[0] != plus {
			num *= -1
		}
		nums = append(nums, num)
	}
	return nums
}

func getSum(src []int) int {
	sum := 0
	for _, v := range src {
		sum += v
	}
	return sum
}

func firstChallenge(src []int) {
	fmt.Println(getSum(src))
}

func getDupF(src []int) int {
	memo, cur, pointer := map[int]bool{0: true}, 0, 0
	for {
		cur, pointer = cur+src[pointer], (pointer+1)%len(src)
		if memo[cur] {
			return cur
		}
		memo[cur] = true
	}
}

func secondChallenge(src []int) {
	fmt.Println(getDupF(src))
}

func main() {
	input := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(input)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(input)
}
