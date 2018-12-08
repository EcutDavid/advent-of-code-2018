package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	nums := []int{}
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}

type node struct {
	children []*node
	metadata []int
}

func buildTree(nums []int, count, l int, parent *node) int {
	pointer := l
	for count > 0 {
		count--
		metaLen, newNode := nums[pointer+1], node{
			[]*node{},
			[]int{},
		}
		parent.children = append(parent.children, &newNode)

		pointer = buildTree(nums, nums[pointer], pointer+2, &newNode)
		for i := 0; i < metaLen; i++ {
			newNode.metadata = append(newNode.metadata, nums[pointer+i])
		}
		pointer += metaLen
	}
	return pointer
}

func walkTree(root *node) {
	fmt.Println(root.metadata)
	for _, child := range root.children {
		walkTree(child)
	}
}

func build(nums []int) *node {
	root := node{
		[]*node{},
		[]int{},
	}
	r := len(nums) - nums[1] - 1
	buildTree(nums, nums[0], 2, &root)
	for i := r + 1; i < len(nums); i++ {
		root.metadata = append(root.metadata, nums[i])
	}
	return &root
}

func walk(root *node) int {
	sum := 0
	for _, v := range root.metadata {
		sum += v
	}
	for _, child := range root.children {
		sum += walk(child)
	}
	return sum
}

func getNodeValue(root *node) int {
	sum := 0
	if len(root.children) == 0 {
		for _, v := range root.metadata {
			sum += v
		}
		return sum
	}
	for _, v := range root.metadata {
		if (v > len(root.children)) || (v == 0) {
			continue
		}
		sum += getNodeValue(root.children[v-1])
	}

	return sum
}

func firstChallenge() {
	nums := getInput()
	root := build(nums)

	fmt.Println(walk(root))
}

func secondChallenge() {
	nums := getInput()
	root := build(nums)

	fmt.Println(getNodeValue(root))
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
