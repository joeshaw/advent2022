package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack []string

// ParseInput parses the input file and returns a slice of stacks
// representing the starting positions of the crates, and a slice of strings
// representing the move operations.
func ParseInput(file string) ([]stack, []string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var stacks []stack
	var moves []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == '[' {
			// Parse the starting positions of the crates
			var stack stack
			for _, ch := range line {
				if ch == '[' || ch == ']' {
					continue
				}
				stack = append(stack, string(ch))
			}
			stacks = append(stacks, stack)
		} else {
			// Parse the move operations
			moves = append(moves, line)
		}
	}

	return stacks, moves
}

// Move performs the given move operation on the stacks. It takes the move
// string, the slice of stacks, and the number of stacks as arguments, and
// returns the updated stacks.
func Move(m string, stacks []stack, n int) []stack {
	var from, to int
	fmt.Sscanf(m, "move %d from %d to %d", &from, &from, &to)
	from--
	to--
	crate := stacks[from][0]
	stacks[from] = stacks[from][1:]
	stacks[to] = append([]string{crate}, stacks[to]...)
	return stacks
}

func main() {
	stacks, moves := ParseInput("input.txt")
	n := len(stacks)
	for _, m := range moves {
		stacks = Move(m, stacks, n)
	}
	var result strings.Builder
	for _, stack := range stacks {
		result.WriteString(stack[0])
	}
	fmt.Println(result.String())
}
