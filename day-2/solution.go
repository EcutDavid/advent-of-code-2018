package main

import (
	"bufio"
	"fmt"
	"os"
)

func firstChallenge() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	twoCounter, threeCounter := 0, 0
	for _, word := range words {
		byteMap := map[byte]int{}
		for i := 0; i < len(word); i++ {
			byteMap[word[i]]++
		}
		for _, v := range byteMap {
			if v == 2 {
				twoCounter++
				break
			}
		}
		for _, v := range byteMap {
			if v == 3 {
				threeCounter++
				break
			}
		}
	}
	fmt.Println(twoCounter, threeCounter)
	fmt.Println(twoCounter * threeCounter)
}

func secondChallenge() {
	// Brute force, O(n ^ 2 * len(word))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	fmt.Println(words)
	for i := 0; i < (len(words) - 1); i++ {
		for j := i + 1; j < len(words); j++ {
			diffIndex, diffCount := -1, 0
			for k := 0; k < len(words[j]); k++ {
				if words[j][k] != words[i][k] {
					diffIndex, diffCount = k, diffCount+1
				}
			}
			fmt.Println(diffCount)
			if diffCount != 1 {
				continue
			}
			// Match
			fmt.Println(words[j][:diffIndex] + words[j][diffIndex+1:])
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
