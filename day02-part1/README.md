# Day 2, part 1

OpenGPT had considerably more difficulty with this task.  I could not get it to produce a correct solution with the Advent of Code prompt alone.  It really struggled with scoring win/loss/draw correctly plus the score for the type of position played.  It usually got one right, but not the other.  I had to augment the prompt with clarifications at the end.

Although I liked the match-based solutions it generated in Rust, the first correct program it generated was in Go.

As with Day 1, part 2 I continued the existing OpenGPT conversation so it could use context from previous solutions.

Run the generated code with `go run rock_paper_scissors.go < input.txt`.

## Prompt

We'll move onto the next day's puzzle.  As before, generate a program that will solve this puzzle given the prompt below.  You may reuse and modify your code from the previous answer to solve it.

The prompt follows in the code-fenced section below:

```
The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant Rock Paper Scissors tournament is already in progress.

Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape. Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.

Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.

The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.

The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

Since you can't be sure if the Elf is trying to help you or trick you, you should calculate the score you would get if you were to follow the strategy guide.

For example, suppose you were given the following strategy guide:

A Y
B X
C Z

This strategy guide predicts and recommends the following:

- In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
- In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
- The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

What would your total score be if everything goes exactly according to your strategy guide?
```

Remember:
* Each line represents a round
* The format of a round is an A, B, or C in the first column, a space in the second column, and an X, Y, or Z in the third column
* A, B, C and X, Y, Z represent Rock, Paper, Scissors respectively
* You score one point for playing Rock, 2 for Paper, and 3 for Scissors
* In addition, you score 6 points for winning a round, 3 points for a draw, and 0 for a loss.

## OpenGPT response

Here is a possible solution to the problem in Go:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var score int

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()

        // Split the line into the opponent's move and your move
        moves := strings.Split(line, " ")
        opponent := moves[0]
        player := moves[1]

        // Calculate the outcome of the round
        // (1 = loss, 3 = draw, 6 = win)
        outcome := 0
        if opponent == "A" && player == "Y" ||
            opponent == "B" && player == "Z" ||
            opponent == "C" && player == "X" {
            outcome = 6
        } else if opponent == player {
            outcome = 3
        }

        // Calculate your score for the round
        // (1 = Rock, 2 = Paper, 3 = Scissors)
        score += 0
        if player == "X" {
            score += 1
        } else if player == "Y" {
            score += 2
        } else if player == "Z" {
            score += 3
        }

        // Add the outcome of the round to your score
        score += outcome
    }

    fmt.Println(score)
}
```

This program uses a `Scanner` from the `bufio` package to read the input from standard input. The input consists of a series of lines, each representing a round in the Rock Paper Scissors tournament. The first and second columns on each line represent the opponent's move and your move, respectively.

The program calculates the outcome of each round based on the moves of the opponent and the player. If the player wins, the outcome is 6 points. If the round is a draw, the outcome is 3 points. Otherwise, the outcome is 0 points.

Next, the program calculates the score for the player based on their move. If the player chooses Rock, their score is 1 point. If they choose Paper, their score is 2 points. If they choose Scissors, their score is 3 points.

Finally, the program adds the outcome of the round to the player's score, and repeats this process for each round in the tournament. At the end, it prints the total score for the player.

You can run this program by saving the code to a file (e.g. `rock_paper_scissors.go`), and running the following command:

```
go run rock_paper_scissors.go
```

You can then provide the input to the program through standard input, for example:

```
A Y
B X
C Z
```

This should produce the output `15`, indicating that the player's total score is 15 points.
