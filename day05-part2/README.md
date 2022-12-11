# Day 5, part 2

As I mentioned in part 1, ChatGPT got stuck and I had to start a new conversation for part 2.  After a few failed attempts at giving it the code from part 1 and the full prompts, and eager to move on, I decided to try a different approach.

Without any context, I gave it the program from part 1.  It responded with a description of what the code did.  I then asked it to make the main modification from part 2.  I did not give it the prompt at all.  Because this was a simple modification, it was able to easily generate a correct program from this.

Run the generated code with `go run crates.go`.

## Prompt

Here is the (unused) part 2 prompt:

```
--- Part Two ---

As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.

Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

The CrateMover 9001 is notable for many new and exciting features: air conditioning, leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.

Again considering the example above, the crates begin in the same configuration:

    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 
Moving a single crate from stack 2 to stack 1 behaves the same as before:

[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 
However, the action of moving three crates from stack 1 to stack 3 means that those three moved crates stay in the same order, resulting in this new configuration:

        [D]
        [N]
    [C] [Z]
    [M] [P]
 1   2   3
Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

        [D]
        [N]
[C]     [Z]
[M]     [P]
 1   2   3
Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

        [D]
        [N]
        [Z]
[M] [C] [P]
 1   2   3
In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.

Before the rearrangement process finishes, update your simulation so that the Elves know where they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?
```

## Conversation

### Me

I have this code: 

```
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
```

### ChatGPT

This is a program written in Go that simulates a drawing of stacks of boxes. The program reads a text file named "input.txt" and parses it to extract the instructions for drawing the stacks and for moving boxes from one stack to another. The program then simulates the instructions and prints the top box of each stack.

The `Stack` type is a simple stack data structure implemented using a slice. The `parseDrawing` function processes the input and creates a slice of `Stack`s representing the initial state of the drawing. The `main` function reads the input file and parses the instructions. It then simulates the instructions by moving boxes between the stacks and prints the top box of each stack at the end.

### Me

Can you modify the program so that when multiple boxes are moved between stacks, they maintain their original order?

### ChatGPT

Sure, we can modify the `main` function to maintain the original order of the boxes when moving them between stacks. Here's a modified version of the `main` function that does that:

```go
func main() {
	// ...

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		fromStack := stacks[from-1]
		toStack := stacks[to-1]

		// Move the specified number of boxes from the "from" stack to the "to" stack,
		// maintaining their original order.
		boxes := make([]string, 0, count)
		for i := 0; i < count; i++ {
			box := fromStack.Pop()
			boxes = append(boxes, box)
		}
		for i := len(boxes) - 1; i >= 0; i-- {
			toStack.Push(boxes[i])
		}
	}

	// ...
}
```

This modified version of the `main` function moves the specified number of boxes from the "from" stack to the "to" stack, maintaining their original order. It does this by first popping the boxes from the "from" stack and appending them to a temporary slice of boxes. Then it pops the boxes from the temporary slice and pushes them onto the "to" stack, which maintains their original order.
