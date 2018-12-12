package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pot struct {
	hasPlant bool
	id       int
}
type rule struct {
	pattern []bool
	result  bool
}

func getInput() ([]pot, []rule) {
	p, rules := []pot{}, []rule{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pS := ""
	fmt.Sscanf(scanner.Text(), "initial state: %s", &pS)
	for i := 0; i < len(pS); i++ {
		if pS[i:i+1] == "#" {
			p = append(p, pot{true, i})
		} else {
			p = append(p, pot{false, i})
		}
	}

	scanner.Scan()
	for scanner.Scan() {
		raw, r := strings.Split(scanner.Text(), " => "), rule{}
		if raw[1] == "#" {
			r.result = true
		}
		r.pattern = []bool{}
		for i := 0; i < len(raw[0]); i++ {
			r.pattern = append(r.pattern, raw[0][i:i+1] == "#")
		}
		rules = append(rules, r)
	}
	return p, rules
}

func applyRules(src []pot, r []rule) {
	srcCopy := make([]pot, len(src))
	copy(srcCopy, src)
	for i := 2; i < len(src)-2; i++ {
		appliedRule := false

		for _, ruleItem := range r {
			match := true
			for j := 0; j < len(ruleItem.pattern); j++ {
				if src[i-2+j].hasPlant != ruleItem.pattern[j] {
					match = false
				}
			}
			if match {
				appliedRule, srcCopy[i].hasPlant = true, ruleItem.result
			}
		}

		if !appliedRule {
			srcCopy[i].hasPlant = false
		}
	}
	for i := 0; i < len(src); i++ {
		src[i].hasPlant = srcCopy[i].hasPlant
	}
}

func inspectPots(p []pot) {
	for _, v := range p {
		if v.hasPlant {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
}

func getPlantSum(p []pot) {
	sum := 0
	for _, v := range p {
		if v.hasPlant {
			sum++
		}
	}
	fmt.Print(sum, "*")
}

func getAnswer(p []pot) {
	sum := 0
	for _, v := range p {
		if v.hasPlant {
			sum += v.id
		}
	}
	fmt.Print(sum, "*")
}

func firstChallenge() {
	p, r := getInput()
	dayCount, maxExtend := 20, 4
	pre, after := make([]pot, dayCount*maxExtend), make([]pot, dayCount*maxExtend)
	for i := 0; i < dayCount*maxExtend; i++ {
		pre[i].id = -dayCount*maxExtend + i
		after[i].id = len(p) + i
	}
	p = append(append(pre, p...), after...)
	for i := 0; i < 20; i++ {
		applyRules(p, r)
	}
	sum := 0
	for _, v := range p {
		if v.hasPlant {
			sum += v.id
		}
	}
	fmt.Println(sum)
}

// "Solved" the problem by looking at console output, and finding the pattern.... Bad....
// After certain loops, the pots display a "pattern" that I have confidence that gonna be kept even loop count become very large.
// If the "pattern" appear, just check what's the increment each loop brings, and calc the final result.
func secondChallenge() {
	p, r := getInput()
	dayCount, maxExtend := 30, 4
	pre, after := make([]pot, dayCount*maxExtend), make([]pot, dayCount*maxExtend)
	for i := 0; i < dayCount*maxExtend; i++ {
		pre[i].id = -dayCount*maxExtend + i
		after[i].id = len(p) + i
	}
	p = append(append(pre, p...), after...)
	getPlantSum(p)
	for i := 0; i < 200; i++ {
		applyRules(p, r)
		inspectPots(p)
		fmt.Print("", i+1, "%%%")
		getAnswer(p)
		fmt.Println()
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}
