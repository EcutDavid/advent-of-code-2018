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

func firstChallenge(IP int, ops []op) {
	state, opLen := [6]int64{}, int64(len(ops))
	for state[IP] < opLen {
		op := ops[state[IP]]
		state = funcMap[op.name](state, op.a, op.b, op.c)
		(&state)[IP]++
	}
	fmt.Println(state[0])
}

func secondChallenge(IP int, ops []op) {
	// The Code can be translated to below with my test case.

	// s := [6]int64{0, 3, 1, 0, 10551261, 1}
	// for {
	// 	if s[4] == s[5]*s[2] {
	// 		(&s)[0] = s[0] + s[5]
	// 	}
	// 	(&s)[2]++
	// 	if s[2] > s[4] {
	// 		(&s)[5]++
	// 		// fmt.Println(s[5])
	// 		if s[5] > s[4] {
	// 			fmt.Println(s[0])
	// 			os.Exit(0)
	// 		}
	// 		(&s)[2] = 1
	// 	}
	// }

	// Code above can be translated to below.
	sum := 0
	for i := 1; i <= 10551261; i++ {
		if 10551261%i == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

func main() {
	IP, ops := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(IP, ops)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(IP, ops)
}
