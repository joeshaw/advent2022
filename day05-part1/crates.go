package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() string {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.items)
}

func printStacks(stacks []*Stack) {
	for _, stack := range stacks {
		fmt.Println(stack)
	}
}

func parseDrawing(lines []string) []*Stack {
	var stacks []*Stack
	for i := 0; i < len(lines[0]); i++ {
		stack := &Stack{}
		for j := 0; j < len(lines); j++ {
			if lines[j][i] != '[' && lines[j][i] != ']' && lines[j][i] != ' ' {
				stack.items = append([]string{string(lines[j][i])}, stack.items...)
			}
		}
		if len(stack.items) > 0 {
			stacks = append(stacks, stack)
		}
	}
	return stacks
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	var instructions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "move") {
			instructions = append(instructions, line)
			continue
		}
		lines = append(lines, line)
	}

	stacks := parseDrawing(lines[:len(lines)-1])
	printStacks(stacks)

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		fromStack := stacks[from-1]
		toStack := stacks[to-1]
		for i := 0; i < count; i++ {
			item := fromStack.Pop()
			toStack.Push(item)
		}
	}

	for _, stack := range stacks {
		fmt.Println(stack.Peek())
	}
}
