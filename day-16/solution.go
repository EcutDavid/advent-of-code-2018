package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkIndex(i int) bool {
	if i > 3 {
		return false
	}
	return true
}

func addr(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(b) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a] + r[b]

	return newR
}

func addi(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a] + b

	return newR
}

func mulr(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(b) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a] * r[b]

	return newR
}

func muli(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a] * b

	return newR
}

func banr(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(b) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a] & r[b]

	return newR
}

func bani(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a] & b

	return newR
}

func borr(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(b) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a] | r[b]

	return newR
}

func bori(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a] | b

	return newR
}

func setr(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = r[a]

	return newR
}

func seti(r [4]int, a, b, c int) [4]int {
	if !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	(&newR)[c] = a

	return newR
}

func gtir(r [4]int, a, b, c int) [4]int {
	if !checkIndex(b) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	if a > r[b] {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}

	return newR
}

func gtri(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	if r[a] > b {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}

	return newR
}

func gtrr(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(b) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	if r[a] > r[b] {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}

	return newR
}

func eqir(r [4]int, a, b, c int) [4]int {
	if !checkIndex(b) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	if a == r[b] {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}

	return newR
}

func eqri(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	if r[a] == b {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}

	return newR
}

func eqrr(r [4]int, a, b, c int) [4]int {
	if !checkIndex(a) || !checkIndex(b) || !checkIndex(c) {
		return [4]int{-1, -1, -1, -1}
	}
	newR := r
	if r[a] == r[b] {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}

	return newR
}

var funcMap = map[string]func(r [4]int, a, b, c int) [4]int{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

func firstChallenge(ops []op) {
	sum := 0
	for _, o := range ops {
		opSum := 0
		for _, v := range funcMap {
			if v(o.src, o.cmd[1], o.cmd[2], o.cmd[3]) == o.res {
				opSum++
			}
		}
		if opSum >= 3 {
			sum++
		}
	}
	fmt.Println("Total OPS that has three or more op match", sum)
}

func secondChallenge(ops []op, realCmds [][4]int) {
	opsMap := map[int]map[string]bool{}
	for _, o := range ops {
		isNew := false
		if opsMap[o.cmd[0]] == nil {
			isNew = true
		}
		newMap := map[string]bool{}
		for k, v := range funcMap {
			if v(o.src, o.cmd[1], o.cmd[2], o.cmd[3]) == o.res {
				newMap[k] = true
			}
		}
		if isNew {
			opsMap[o.cmd[0]] = newMap
			continue
		}
		for k, v := range opsMap[o.cmd[0]] {
			if !v {
				continue
			}
			// The same op ID does not work for the newMap, means a contradition
			if !newMap[k] {
				opsMap[o.cmd[0]][k] = false
			}
		}
	}
	idOpMap, done := map[int]string{}, map[string]bool{}
	for len(idOpMap) < 16 {
		for k, v := range opsMap {
			_, ok := idOpMap[k]
			// This one already handled
			if ok {
				continue
			}
			sum, lastOpId := 0, ""
			for opId, v := range v {
				if v && !done[opId] {
					sum, lastOpId = sum+1, opId
				}
			}
			if sum == 1 {
				idOpMap[k], done[lastOpId] = lastOpId, true
			}
		}
	}
	initState := [4]int{0, 0, 0, 0}
	for _, v := range realCmds {
		initState = funcMap[idOpMap[v[0]]](initState, v[1], v[2], v[3])
	}
	fmt.Println("final state", initState)
}

type op struct {
	src, res [4]int
	cmd      [4]int
}

func parseInput() ([]op, [][4]int) {
	scanner := bufio.NewScanner(os.Stdin)
	res, realCmds := []op{}, [][4]int{}
	processOp, index, a, b, c, d := false, 0, 0, 0, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 6 {
			continue
		}
		if line[:6] == "Before" {
			fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &a, &b, &c, &d)
			newOp := op{}
			newOp.src = [4]int{a, b, c, d}
			res, processOp = append(res, newOp), true
		} else if line[:5] == "After" {
			fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &a, &b, &c, &d)
			res[index].res = [4]int{a, b, c, d}
			processOp, index = false, index+1
		} else {
			fmt.Sscanf(line, "%d %d %d %d", &a, &b, &c, &d)
			if processOp {
				res[index].cmd = [4]int{a, b, c, d}
			} else {
				realCmds = append(realCmds, [4]int{a, b, c, d})
			}

		}
	}
	return res, realCmds
}

func main() {
	ops, realCmds := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(ops)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(ops, realCmds)
}
