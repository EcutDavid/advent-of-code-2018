package main

import (
	"fmt"
)

type dobuleLinkedList struct {
	next *dobuleLinkedList
	prev *dobuleLinkedList
	val  int
}

func (l *dobuleLinkedList) insert(val int) *dobuleLinkedList {
	newNode := dobuleLinkedList{val: val}
	oldNext := l.next
	l.next, oldNext.prev = &newNode, &newNode
	newNode.prev, newNode.next = l, oldNext
	return &newNode
}
func (l *dobuleLinkedList) delete() *dobuleLinkedList {
	res := l.next
	l.prev.next, l.next.prev = res, res
	return res
}

func betterClac(n, p int) {
	score, node := make([]int64, p), &dobuleLinkedList{val: 0}
	node.next, node.prev = node, node

	counter := 1
	for i := 1; i <= p; i++ {
		if i%23 == 0 {
			for j := 0; j < 7; j++ {
				node = node.prev
			}
			score[counter] += int64(i + node.val)
			node = node.delete()
		} else {
			node = node.next.insert(i)
		}
		counter = (counter + 1) % n
	}

	max := int64(0)
	for i := 0; i < len(score); i++ {
		if score[i] > max {
			max = score[i]
		}
	}
	fmt.Println(max)
}

func main() {
	betterClac(9, 25)
	betterClac(10, 1618)
	betterClac(13, 7999)
	betterClac(17, 1104)
	betterClac(21, 6111)
	betterClac(30, 5807)
	betterClac(403, 71920)
	betterClac(403, 7192000)
}
