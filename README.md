# Advent of Code 2018
My solutions for [adventofcode.com/2018](https://adventofcode.com/2018)

## My Solutions
* Day 1: For problem 2, creating a hashtable for the F already being created, whenever there is a new one collide with existing one, done.
* Day 2: For problem 2, A brute force O(n ^ 2 * len(word)) solution that for each line, check all the line after it, whether the "diff" counter is equal to 1.
* Day 3: Because the "claims" all have small width and height, simply create a two-dimensional array, initialize all the slot to 0, and increment the slot based on the claim, the answer can be found.
* Day 4: Sort the input lines alphabetically, the result would be time increasing order, the answer then can be found by analysing the intersections.
* Day 5: Using a doubly linked list to sort all the items, loop them until there is a loop has no collision, whenever there is a collision, drop two nodes from the list.
* Day 15: A fun breadth-first graph search question, very easy to miss conditions and edge cases, a super nice graph question!