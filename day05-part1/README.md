# Day 5, part 1

Hoo boy.  I thought I was going to give up on this one.  ChatGPT really struggled to understand the columnar nature of the input.  I tried a lot of different tactics but the one that ultimately worked is below.

This one was a lot more conversational and interative.  Once I was able to get it to output a reasonable program and pivot it to understanding columns, I was able to slowly get it to a program that gave the right answer.

The code is really ugly, and not at all how I would have written it.  But it works.

The output was one letter per line:
```
P
S
N
R
G
B
T
F
T
```

It proved very difficult to get it to output this on one line!  And eventually ChatGPT locked up on me and I couldn't do it.  But at least I had gotten the right answer (if not the right format) before that happened.

Given that I lost my conversation state, though, I am not sure how I am going to solve part 2, which will be a variation on this problem.  Hopefully pasting in the generated code will be enough!

Run the generated code with `go run crates.go`.

Below is the full conversation.  I've added block quotes inline for additional commentary.

## Conversation

### Me

Advent of Code is a Christmas-themed series of programming puzzles. Each day there are two parts to a puzzle, with the second part building upon the first one. I will provide you with the prompt for the first part. I want you to generate a complete program in Go, Rust, or Python to solve the puzzle. The program should read its input from a file named `input.txt`.

The prompt will follow below.  I'll put code and ASCII diagrams in triple backtick fenced blocks.

> In earlier attempts I felt that ChatGPT struggled with the formatting of the prompt.  I modified the prompt to code fence the bits that were input to be processed.

The prompt:
---
The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

```
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
```

In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

```
[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 
```

In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

```
        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3
```

Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

```
        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3
```

Finally, one crate is moved from stack 1 to stack 2:

```
        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3
```

The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each stack?
---

Ignore the example drawings from the prompt.  Use this sample input instead:
```
[B]                     [N]     [H]
[V]         [P] [T]     [V]     [P]
[W]     [C] [T] [S]     [H]     [N]
[T]     [J] [Z] [M] [N] [F]     [L]
[Q]     [W] [N] [J] [T] [Q] [R] [B]
[N] [B] [Q] [R] [V] [F] [D] [F] [M]
[H] [W] [S] [J] [P] [W] [L] [P] [S]
[D] [D] [T] [F] [G] [B] [B] [H] [Z]
 1   2   3   4   5   6   7   8   9

move 2 from 8 to 1
move 4 from 9 to 8
move 2 from 1 to 6
move 7 from 4 to 2
move 10 from 2 to 7
move 2 from 1 to 6
```

> I needed to tell ChatGPT to "ignore the example drawings from the prompt" because it overfitted them, and wrote programs that were hardcoded to handle inputs with three columns.

There are two parts to the input file.  The first part is the stack drawing.  The second part is the rearrangement procedure.  They are separated by a blank line.

The output of this program should be the top crate of each stack.  If we look at the input of this program, before the rearrangement procedure is run, the values are "BBCPTNNRH".  Print the value after the rearrangement procedure is run.

### ChatGPT

```go
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

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var stacks []*Stack
    var instructions []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "move") {
            instructions = append(instructions, line)
            continue
        }

        items := strings.Split(line, " ")
        stack := &Stack{}
        for _, item := range items {
            if item != "" {
                stack.Push(item)
            }
        }
        stacks = append(stacks, stack)
    }

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
```

### Me

Print the contents of each stack after you have created them but before you run any instructions

> This initial version panicked on the input, so I wanted to see what the stacks looked like that it generated.

### ChatGPT

To print the contents of the stacks before running any instructions, we can add a few lines of code to the main function.

First, we will add a new `printStacks` function that takes a list of stacks as a parameter and prints each stack to the console.

```go
func printStacks(stacks []*Stack) {
	for _, stack := range stacks {
		fmt.Println(stack)
	}
}
```

Then, we will call this function immediately after creating the stacks from the input file.

```go
for scanner.Scan() {
	line := scanner.Text()
	if strings.Contains(line, "move") {
		instructions = append(instructions, line)
		continue
	}

	items := strings.Split(line, " ")
	stack := &Stack{}
	for _, item := range items {
		if item != "" {
			stack.Push(item)
		}
	}
	stacks = append(stacks, stack)
}

printStacks(stacks)
```

Here is the complete updated code:

[elided]

> Each time ChatGPT wanted to print the complete code, but it never completed.  I assume it ran out of CPU time to finish.

### Me

This is close, but you are building stacks row by row rather than column by column.  For instance, your first stack is "BNH" but it should be "BVWTQNHD"

### ChatGPT

To fix this, we need to change how we create the stacks from the input file. Instead of iterating over each line and splitting it by spaces, we will iterate over each character in each line and append it to the appropriate stack.

First, we will create a `parseStacks` function that takes a list of lines as a parameter and returns a list of stacks. This function will iterate over each line and each character in the line, and append the character to the appropriate stack.

```go
func parseStacks(lines []string) []*Stack {
	var stacks []*Stack
	for i := 0; i < len(lines[0]); i++ {
		stack := &Stack{}
		for j := 0; j < len(lines); j++ {
			if lines[j][i] != ' ' {
				stack.Push(string(lines[j][i]))
			}
		}
		stacks = append(stacks, stack)
	}
	return stacks
}
```

Then, we will update the main function to use this function to parse the input lines and create the stacks.

```go
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
		if strings.Contains(line, "move") {
			instructions = append(instructions, line)
			continue
		}
		lines = append(lines, line)
	}

	stacks := parseStacks(lines)
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
```

Here is the complete updated code:

[elided]

### Me

This doesn't account for the blank line that separates the drawing from the instructions.  I think it would be better if you handled those in separate `for` loops.

### ChatGPT

To fix this issue, we can update the main function to have two separate for loops, one for parsing the drawing and one for parsing the instructions.

First, we will move the code for parsing the drawing into a separate function called `parseDrawing`. This function will take a list of lines as a parameter and return a list of stacks.

```go
func parseDrawing(lines []string) []*Stack {
	var stacks []*Stack
	for i := 0; i < len(lines[0]); i++ {
		stack := &Stack{}
		for j := 0; j < len(lines); j++ {
			if lines[j][i] != ' ' {
				stack.Push(string(lines[j][i]))
			}
		}
		stacks = append(stacks, stack)
	}
	return stacks
}
```
> `parseDrawing` replaces the previous `parseStacks` function.

Next, we will update the main function to use this function to parse the drawing lines and create the stacks.

```go
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

	stacks := parseDrawing(lines)
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
```

Here is the complete updated code:

[elided]

### Me

You are getting closer!  You should skip over square brackets when parsing the drawing, and be aware that instead of "[X]" there could be three spaces ("   ") denoting a blank space that shouldn't be pushed onto the stack.  Finally, you should disregard the final line of the drawing that just numbers the stacks.

### ChatGPT

To fix these issues, we need to update the `parseDrawing` function.

First, we will change the inner for loop to only append characters that are not square brackets or spaces to the stack.

```go
func parseDrawing(lines []string) []*Stack {
	var stacks []*Stack
	for i := 0; i < len(lines[0]); i++ {
		stack := &Stack{}
		for j := 0; j < len(lines)-1; j++ {
			if lines[j][i] != '[' && lines[j][i] != ']' && lines[j][i] != ' ' {
				stack.Push(string(lines[j][i]))
			}
		}
		stacks = append(stacks, stack)
	}
	return stacks
}
```

Next, we will update the main function to exclude the last line of the input when parsing the drawing.

```go
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
```

### Me

Getting closer.  Don't push empty stacks onto the list.

### ChatGPT

To fix this issue, we need to update the `parseDrawing` function to only append non-empty stacks to the list.

First, we will add a check at the end of the inner for loop to only append the stack to the list if it contains items.

```go
func parseDrawing(lines []string) []*Stack {
	var stacks []*Stack
	for i := 0; i < len(lines[0]); i++ {
		stack := &Stack{}
		for j := 0; j < len(lines)-1; j++ {
			if lines[j][i] != '[' && lines[j][i] != ']' && lines[j][i] != ' ' {
				stack.Push(string(lines[j][i]))
			}
		}
		if len(stack.items) > 0 {
			stacks = append(stacks, stack)
		}
	}
	return stacks
}
```

Here is the complete updated code:

[elided]

### Me

Based on your stack implementation, the stacks are backward and need to be reversed.

### ChatGPT

To fix this issue, we need to update the `parseDrawing` function to reverse the order of the items in the stack.

First, we will update the inner for loop to push items onto the front of the stack rather than the back.

```go
func parseDrawing(lines []string) []*Stack {
	var stacks []*Stack
	for i := 0; i < len(lines[0]); i++ {
		stack := &Stack{}
		for j := 0; j < len(lines)-1; j++ {
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
```

Here is the complete updated code:

[elided]

### Me

I just noticed that this is skipping the last line of the drawing input.  I think you need to make it `j < len(lines)`, not `j < len(lines)-1`.

> I later noticed this bug was introduced in an earlier part of the conversation where I told ChatGPT to "disregard the final line of the drawing that just numbers the stacks."

### ChatGPT

To fix this issue, we need to update the `parseDrawing` function to include the last line of the input when parsing the drawing.

First, we will change the for loop condition to `j < len(lines)` rather than `j < len(lines)-1`.

```go
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
```

Here is the complete updated code:

[elided]

### Me

When walking across the stacks at the end of the main function, print the output on a single line.  For example, "PSNRGBTFT".

> ChatGPT generated wrong code in response to this, trying to update the `printStacks` function rather than the simple loop at the end of `main`.  Unfortunately, ChatGPT got wedged at this point and I couldn't make further progress without losing state.
>
> That's ok though, because at this point it printed the right answer, just not in the exact format I wanted.
