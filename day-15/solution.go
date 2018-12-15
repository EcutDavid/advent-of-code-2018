// A great DFS problem!
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const (
	wall = iota
	empty
)
const elf = uint(2)
const goblin = uint(4)
const attackPower = 3
const defaultHP = 200

var wallASC = "#"[0]
var elfASC = "E"[0]
var goblinASC = "G"[0]

type unit struct {
	die   bool
	pos   [2]int
	hp    int
	isElf bool
}
type board [][]int
type unitList []unit

func (u unitList) Len() int      { return len(u) }
func (u unitList) Swap(i, j int) { u[i], u[j] = u[j], u[i] }
func (u unitList) Less(i, j int) bool {
	if u[i].pos[0] == u[j].pos[0] {
		return u[i].pos[1] < u[j].pos[1]
	}
	return u[i].pos[0] < u[j].pos[0]
}

// Only for DEBUG
func draw(board board, units unitList) {
	unitMap := genUnitMap(units)
	for y, line := range board {
		for x, slot := range line {
			v, ok := unitMap[[2]int{y, x}]
			if ok {
				if v.isElf {
					fmt.Print("E")
				} else {
					fmt.Print("G")
				}
				continue
			}
			if slot == wall {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// Only for DEBUG
func reportHP(units unitList) {
	for _, v := range units {
		if v.die {
			continue
		}
		fmt.Println("isElf", v.isElf, "line:", v.pos[0]+1, "hp:", v.hp)
	}
}

func parseInput() (board, unitList) {
	scanner := bufio.NewScanner(os.Stdin)
	board, y, unites := [][]int{}, 0, unitList{}

	for scanner.Scan() {
		text := scanner.Text()
		line := make([]int, len(text))
		for i := 0; i < len(text); i++ {
			if text[i] == wallASC {
				line[i] = wall
				continue
			}
			line[i] = empty
			if text[i] == elfASC {
				unites = append(unites, unit{false, [2]int{y, i}, defaultHP, true})
			} else if text[i] == goblinASC {
				unites = append(unites, unit{false, [2]int{y, i}, defaultHP, false})
			}
		}
		board, y = append(board, line), y+1
	}
	return board, unites
}

func genUnitMap(units unitList) map[[2]int]*unit {
	res := map[[2]int]*unit{}
	for i := 0; i < len(units); i++ {
		if units[i].die {
			continue
		}
		res[units[i].pos] = &units[i]
	}
	return res
}

// Answers what kind of unit can start attack
func genAttackMap(units unitList) map[[2]int]uint {
	res := map[[2]int]uint{}
	for _, v := range units {
		if v.die {
			continue
		}
		slot := elf
		if v.isElf {
			slot = goblin
		}
		res[[2]int{v.pos[0] + 1, v.pos[1]}] |= slot
		res[[2]int{v.pos[0] - 1, v.pos[1]}] |= slot
		res[[2]int{v.pos[0], v.pos[1] + 1}] |= slot
		res[[2]int{v.pos[0], v.pos[1] - 1}] |= slot
	}
	return res
}

var dirs = [4][2]int{
	[2]int{-1, 0},
	[2]int{0, -1},
	[2]int{0, 1},
	[2]int{1, 0},
}

func canAttack(u unit, pos [2]int, attackMap map[[2]int]uint) bool {
	if u.isElf {
		return (attackMap[pos] & elf) > 0
	}
	return (attackMap[pos] & goblin) > 0
}

// Return the next move, if next move is current condition, means can attack.
// If returns not okay, then, cannot move, skip
func bfs(board board, u unit, attackMap map[[2]int]uint, unitMap map[[2]int]*unit) (bool, [2]int) {
	_, okay := attackMap[u.pos]
	if okay && canAttack(u, u.pos, attackMap) {
		return true, u.pos
	}

	visited, parent := map[[2]int]bool{u.pos: true}, map[[2]int][2]int{}
	queue, distance := [][2]int{u.pos}, map[[2]int]int{u.pos: 0}
	decisions := [][2]int{}
	for len(queue) > 0 {
		task := queue[0]
		_, okay := attackMap[task]
		if okay && canAttack(u, task, attackMap) {
			taskCopy := task
			for parent[taskCopy] != u.pos {
				taskCopy = parent[taskCopy]
			}
			return true, taskCopy
		}
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			newPos := [2]int{task[0] + dirs[i][0], task[1] + dirs[i][1]}
			if (newPos[0] < 0) || (newPos[0] >= len(board)) || (newPos[1] < 0) || (newPos[1] >= len(board[0])) {
				continue
			}
			_, ok := unitMap[newPos]
			if ok || (board[newPos[0]][newPos[1]] == wall) || visited[newPos] {
				continue
			}
			visited[newPos], queue, parent[newPos] = true, append(queue, newPos), task
			distance[newPos] = distance[task] + 1
		}
	}
	if len(decisions) > 0 {
		return true, decisions[0]
	}
	return false, [2]int{}
}

func attack(u unit, unitMap map[[2]int]*unit, power int) {
	decisions := [][2]int{}
	for _, v := range dirs {
		newPos := [2]int{u.pos[0] + v[0], u.pos[1] + v[1]}
		target, ok := unitMap[newPos]
		if ok && target.isElf != u.isElf {
			decisions = append(decisions, newPos)
		}
	}
	if len(decisions) == 0 {
		return
	}

	sameHP, lastHP, minHp, minHpIndex := true, unitMap[decisions[0]].hp, math.MaxInt32, 0
	for k, v := range decisions {
		if minHp > unitMap[v].hp {
			minHpIndex, minHp = k, unitMap[v].hp
		}
		if unitMap[v].hp != lastHP {
			sameHP = false
		}
	}
	target := unitMap[decisions[0]]
	if !sameHP {
		target = unitMap[decisions[minHpIndex]]
	}
	target.hp -= power
	if target.hp <= 0 {
		target.die = true
	}
}

func isDone(units unitList) bool {
	elfSum, gSum := 0, 0
	for _, v := range units {
		if v.die {
			continue
		}
		if v.isElf {
			if gSum > 0 {
				return false
			}
			elfSum++
		} else {
			if elfSum > 0 {
				return false
			}
			gSum++
		}
	}
	return true
}

func getScore(round int, units unitList) (sum int) {
	for _, v := range units {
		if v.die {
			continue
		}
		sum += v.hp
	}
	return sum * round
}

func simulate(board board, units unitList, elfPower int) {
	// 5... just a big number choosed randomly
	for i := 0; i < 5000000; i++ {
		sort.Sort(units)
		done, doneEarly := false, true
		for i := 0; i < len(units); i++ {
			if units[i].die {
				continue
			}
			unitMap := genUnitMap(units)
			attackMap := genAttackMap(units)
			ok, nextMove := bfs(board, units[i], attackMap, unitMap)
			if !ok {
				continue
			}
			power := attackPower
			if units[i].isElf {
				power = elfPower
			}
			if nextMove == units[i].pos {
				attack(units[i], unitMap, power)
			} else {
				units[i].pos = nextMove
				attack(units[i], unitMap, power)
			}
			if isDone(units) {
				done = true
				if i == (len(units) - 1) {
					doneEarly = false
				}
				break
			}
		}

		if done {
			if doneEarly {
				fmt.Println(getScore(i, units))
			} else {
				fmt.Println(getScore(i+1, units))
			}
			break
		}
	}
}

func firstChallenge(board board, units unitList) {
	simulate(board, units, attackPower)
}

func secondChallenge(board board, units unitList) {
	for i := 3; i <= 200; i++ {
		listCopy := make([]unit, len(units))
		copy(listCopy, units)
		simulate(board, listCopy, i)
		works := true
		for _, v := range listCopy {
			if v.isElf && v.die {
				works = false
			}
		}
		if works {
			os.Exit(0)
		}
	}
}

func main() {
	board, units := parseInput()
	listCopy := make([]unit, len(units))
	copy(listCopy, units)
	fmt.Println("first challenge:")
	firstChallenge(board, units)
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(board, listCopy)
}
