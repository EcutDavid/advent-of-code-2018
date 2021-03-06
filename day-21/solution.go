package main

import (
	"bufio"
	"fmt"
	"os"
)

func addr(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a] + r[b]
	return newR
}

func addi(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a] + b

	return newR
}

func mulr(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a] * r[b]
	return newR
}

func muli(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a] * b
	return newR
}

func banr(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a] & r[b]
	return newR
}

func bani(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a] & b
	return newR
}

func borr(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a] | r[b]
	return newR
}

func bori(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a] | b
	return newR
}

func setr(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = r[a]
	return newR
}

func seti(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	(&newR)[c] = a
	return newR
}

func gtir(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	if a > r[b] {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}
	return newR
}

func gtri(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	if r[a] > b {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}
	return newR
}

func gtrr(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	if r[a] > r[b] {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}
	return newR
}

func eqir(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	if a == r[b] {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}
	return newR
}

func eqri(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	if r[a] == b {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}
	return newR
}

func eqrr(r [6]int64, a, b, c int64) [6]int64 {
	newR := r
	if r[a] == r[b] {
		(&newR)[c] = 1
	} else {
		(&newR)[c] = 0
	}
	return newR
}

var funcMap = map[string]func(r [6]int64, a, b, c int64) [6]int64{
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

type op struct {
	name    string
	a, b, c int64
}

func firstChallenge(IP int, ops []op) {
	state, opLen := [6]int64{}, int64(len(ops))
	for state[IP] < opLen {
		op := ops[state[IP]]
		if op.name == "eqrr" {
			fmt.Println(state)
			return
		}
		state = funcMap[op.name](state, op.a, op.b, op.c)
		(&state)[IP]++
	}
}

func secondChallenge(IP int, ops []op) {
	state, opLen, max, time := [6]int64{}, int64(len(ops)), int64(0), int64(0)
	regFiveMap := map[int64]int64{}
	for state[IP] < opLen {
		op := ops[state[IP]]
		time++
		if op.name == "eqrr" {
			if state[5] > max {
				max = state[5]
			}
			if regFiveMap[state[5]] > 0 {
				timeMax, bestChoice := int64(0), int64(0)
				for k, v := range regFiveMap {
					if v > timeMax {
						timeMax = v
						bestChoice = k
					}
				}
				fmt.Println(bestChoice)
				os.Exit(0)
			}
			regFiveMap[state[5]] = time
		}
		state = funcMap[op.name](state, op.a, op.b, op.c)
		(&state)[IP]++
	}
}

func parseInput() (int, []op) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	IP := 0
	fmt.Sscanf(scanner.Text(), "#ip %d", &IP)
	ops := []op{}
	for scanner.Scan() {
		name, a, b, c := "", int64(0), int64(0), int64(0)
		fmt.Sscanf(scanner.Text(), "%s %d %d %d", &name, &a, &b, &c)
		ops = append(ops, op{name, a, b, c})
	}
	return IP, ops
}

func main() {
	IP, ops := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(IP, ops)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(IP, ops)
}
