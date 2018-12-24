# Advent of Code 2018
My solutions for [adventofcode.com/2018](https://adventofcode.com/2018)

## My Solutions
* Day 1: For the problem 2, creating a hash table for the F already being created, whenever there is a new one collide with existing one, done.
* Day 2: For the problem 2, A brute force O(n ^ 2 * len(word)) solution that for each line, check all the line after it, whether the "diff" counter is equal to 1.
* Day 3: Because the "claims" all have small width and height, simply create a two-dimensional array, initialize all the slot to 0, and increment the slot based on the claim, the answer can be found.
* Day 4: Sort the input lines alphabetically, the result would be time increasing order, the answer can be found by analysing the intersections.
* Day 5: Using a doubly linked list to sort all the items, loop them until there is a loop has no collision, whenever there is a collision, drop two nodes from the list.
* Day 14: Can be solved by brute force, just keep doing what the question asks us to do until the answer appears.
* Day 15: A fun breadth-first graph search question, very easy to miss conditions and edge cases, a super nice graph question!
* Day 16: Need carefully simulate the 16 ASM operations, can use hashmaps to generate the relationship between id(the first number) and operation, due to the input is small, didn't try to make the algorithm run faster.
* Day 18: For the problem 2, instead of brute force to simulate the 1e9 "rounds", should take "snapshot" of each board, once there is two snapshots which already appeared before, the problem can be solved.
* Day 19: This is a day-16 follow up question, problem 1 can be solved with the solution from day-16, problem 2 is very tricky, had decoded the program into a high-level code(or just natural language), after that, we can understand what the program is doing, and optimize the program's time complexity.
* Day 20: Parsing a graph from the input, then, DFS to walk through the map, the hard part of parsing the input, used a recursive solution to solve that part. Wrote a more detailed explanation in [this blog](https://medium.com/@davidguandev/aoc-2018-day-20-a-regular-map-1ef024e85c22).
* Day 21: This is a day-19 follow up question, need to parse the asm code then find the trick there, to solve the problem.
* Day 22: From problem 2, need to generate a map that is bigger than (0, 0) - (`targetX` - `targetY`), and then, run Bellman–Ford or Dijkstra's algorithm to get the shortest path to the target.
* Day 23: Still work in progress :( .
* Day 24: A pretty starigforawrd challenge need use binary search for the second challenge to spend less time on validations.
