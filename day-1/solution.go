package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func firstChallenge() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line[1:])
		if line[0:1] == "+" {
			res += num
		} else {
			res -= num
		}
	}
	fmt.Println(res)
}

func secondChallenge() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	cur, pointer, nums, memo := 0, 0, []int{}, map[int]bool{0: true}

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line[1:])
		if line[0:1] == "+" {
			nums = append(nums, num)
		} else {
			nums = append(nums, -num)
		}
	}

	for {
		cur, pointer = cur+nums[pointer], (pointer+1)%len(nums)
		if !memo[cur] {
			memo[cur] = true
		} else {
			fmt.Println(cur)
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
