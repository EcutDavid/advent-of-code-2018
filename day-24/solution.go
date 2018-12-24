// A pretty starigforawrd challenge, need use binary search for second challenge to spend less time on validations.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const maxRound = 100000

type group struct {
	weaknesses, immunities            map[string]bool
	unitCount, hp, initiative, damage int
	damageType                        string
	isImmune, choosed                 bool
	target, boost                     int
}
type army []*group

func (g *group) getEffective() int {
	return g.unitCount * (g.damage + g.boost)
}
func (g *group) isLive() bool { return g.unitCount > 0 }
func (g *group) getDamage(target *group) int {
	if target.immunities[g.damageType] {
		return 0
	}
	rate := 1
	if target.weaknesses[g.damageType] {
		rate = 2
	}
	return (g.damage + g.boost) * g.unitCount * rate
}
func (g *group) chooseTarget(a army) int {
	mostDamage, index := 0, -1
	for k, d := range a {
		if (d.isImmune == g.isImmune) || !d.isLive() || d.choosed {
			continue
		}
		damage := g.getDamage(d)
		if damage > mostDamage {
			mostDamage, index = damage, k
		}
	}
	if index > -1 {
		a[index].choosed = true
	}
	return index
}

func (a army) Len() int      { return len(a) }
func (a army) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a army) Less(i, j int) bool {
	if a[i].getEffective() == a[j].getEffective() {
		return a[i].initiative > a[j].initiative
	}
	return a[i].getEffective() > a[j].getEffective()
}
func (a army) initRound(immuneBoost int) {
	for _, g := range a {
		g.choosed, g.target = false, -1
		if g.isImmune {
			g.boost = immuneBoost
		}
	}
}
func (a army) copy() army {
	groups, copy := make([]group, len(a)), army{}
	for i := 0; i < len(a); i++ {
		groups[i] = *a[i]
		copy = append(copy, &groups[i])
	}
	return copy
}

func gameFinished(a army) bool {
	s := gameScore(a)
	return (s[0] == 0) || (s[1] == 0)
}
func gameScore(a army) []int {
	s := []int{0, 0}
	for _, g := range a {
		if !g.isLive() {
			continue
		}
		index := 1
		if g.isImmune {
			index = 0
		}
		s[index] += g.unitCount
	}
	return s
}

func simulation(a army, boost int) int {
	roundCounter, initiativeMap, maxInitiative := 0, map[int]*group{}, 0
	for _, g := range a {
		initiativeMap[g.initiative] = g
		if g.initiative > maxInitiative {
			maxInitiative = g.initiative
		}
	}
	for !gameFinished(a) {
		a.initRound(boost)
		sort.Sort(a)
		for _, g := range a {
			if !g.isLive() {
				continue
			}
			g.target = g.chooseTarget(a)
		}
		for i := maxInitiative; i >= 0; i-- {
			attacker := initiativeMap[i]
			if (attacker == nil) || !attacker.isLive() || (attacker.target == -1) {
				continue
			}
			t := a[attacker.target]
			t.unitCount -= attacker.getDamage(t) / t.hp
		}
		roundCounter++
		if roundCounter == maxRound {
			break
		}
	}
	return roundCounter
}

func firstChallenge(a army) {
	simulation(a, 0)
	s := gameScore(a)
	if s[0] == 0 {
		fmt.Println(s[1])
	} else {
		fmt.Println(s[0])
	}
}

func secondChallenge(src army) {
	l, r, best := 0, int(1e9), -1
	for l <= r {
		a, mid := src.copy(), (l+r)/2
		roundCounter := simulation(a, mid)
		s := gameScore(a)
		immuneWin := s[0] != 0
		if roundCounter == maxRound {
			immuneWin = false
		}
		if immuneWin {
			r, best = mid-1, s[0]
		} else {
			l = mid + 1
		}
	}
	fmt.Printf("Immune win with unit count %d\n", best)
}

func parseInput() army {
	result := army{}
	lines, scanner, pointer := make([][]string, 2), bufio.NewScanner(os.Stdin), 0
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "Infection:" {
			pointer++
			continue
		}
		if line == "" {
			continue
		}
		lines[pointer] = append(lines[pointer], line)
	}
	parseGroups := func(lines []string, isImmune bool) []*group {
		groups := []*group{}
		for _, l := range lines {
			u, h, d, i, t, p := 0, 0, 0, 0, "", ""
			pl, pr := strings.Index(l, "("), strings.Index(l, ")")
			if pl != -1 {
				p, l = l[pl+1:pr], l[:pl]+l[pr+1:]
			}
			fmt.Sscanf(
				l,
				"%d units each with %d hit points with an attack that does %d %s damage at initiative %d",
				&u, &h, &d, &t, &i,
			)
			weaknesses, immunities := map[string]bool{}, map[string]bool{}
			if len(p) > 0 {
				for _, s := range strings.Split(p, "; ") {
					if s[0] == "w"[0] {
						for _, d := range strings.Split(s[len("weak to "):], ", ") {
							weaknesses[d] = true
						}
					}
					if s[0] == "i"[0] {
						for _, d := range strings.Split(s[len("immune to "):], ", ") {
							immunities[d] = true
						}
					}
				}
			}
			groups = append(groups, &group{weaknesses, immunities, u, h, i, d, t, isImmune, false, -1, 0})
		}
		return groups
	}

	for _, g := range append(parseGroups(lines[1], false), parseGroups(lines[0], true)...) {
		result = append(result, g)
	}

	return result
}

func main() {
	a := parseInput()
	fmt.Println("first challenge:")
	firstChallenge(a.copy())
	fmt.Println("****************")
	fmt.Println("second challenge:")
	secondChallenge(a)
}
