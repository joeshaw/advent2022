# Day 1, part 1

Run the generated code with `go run calorie_counting.go < input.txt`.

## Prompt

Advent of Code is a Christmas-themed series of programming puzzles.  Each day there are two puzzles that typically build upon the puzzles and answers from the previous ones.
I will provide you with the current puzzle.  I want you to generate a program to solve the puzzle.  You may generate the program in one of the following languages: Go, Rust, or Python.

Today's prompt:

```
--- Day 1: Calorie Counting ---

Santa's reindeer typically eat regular reindeer food, but they need a lot of magical energy to deliver presents on Christmas. For that, their favorite snack is a special type of star fruit that only grows deep in the jungle. The Elves have brought you on their annual expedition to the grove where the fruit grows.

To supply enough magical energy, the expedition needs to retrieve a minimum of fifty stars by December 25th. Although the Elves assure you that the grove has plenty of fruit, you decide to grab any fruit you see along the way, just in case.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
```
## OpenGPT response

Here is a possible solution to the problem in Go:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var maxCalories, totalCalories int

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            if totalCalories > maxCalories {
                maxCalories = totalCalories
            }
            totalCalories = 0
        } else {
            calories, _ := strconv.Atoi(line)
            totalCalories += calories
        }
    }

    fmt.Println(maxCalories)
}
```

This program uses a `Scanner` from the `bufio` package to read the input from standard input. The input consists of a series of numbers representing the number of Calories in each food item. Each Elf's list of items is separated by a blank line.

The program keeps track of the maximum number of Calories seen so far, and the total number of Calories for the current Elf. When it encounters a blank line, it checks if the total number of Calories for the current Elf is greater than the maximum seen so far, and updates the maximum if necessary. Finally, the program prints the maximum number of Calories seen.

You can run this program by saving the code to a file (e.g. `calorie_counting.go`), and running the following command:

```
go run calorie_counting.go
```

You can then provide the input to the program through standard input, for example:

```
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
```

This should produce the output `24000`, indicating that the fourth Elf is carrying the most Calories.
