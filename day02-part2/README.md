# Day 2, part 2

As with part 1, I needed to add additional context to the prompt in order to (eventually) generate a program that gave the correct result.

At first I attempted to reuse the existing conversation, but OpenGPT stopped generating complete programs.  I am not sure if I ran out of compute time or what, but every time it would stop answering in the middle of code generation.

Starting with a new conversation for part 2 was not ideal.  I had to give it both part 1 and part 2 prompts along with additional context.

This prompt generated a complete Go program that solved it correctly and an incomplete Rust program.  It's the first time I've seen it try to solve the problem in multiple languages.

Run the generated code with `go run rock_paper_scissors.go`.

## Prompt

Advent of Code is a Christmas-themed series of programming puzzles.  Each day there are two puzzles, with the second puzzle building upon the first one.  I will provide you with the prompts for both the first and second puzzles.  I want you to generate a program to solve the second puzzle.  It should read its input from a file named `input.txt`.  You may generate the program in one of the following languages: Go, Rust, or Python.

The first prompt:
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

In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

What would your total score be if everything goes exactly according to your strategy guide?

Your puzzle answer was 13009.

The first half of this puzzle is complete! It provides one gold star: *
```

The second prompt:
```
The Elf finishes helping with the tent and sneaks back over to you. "Anyway, the second column says how the round needs to end: X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

The total score is still calculated in the same way, but now you need to figure out what shape to choose so the round ends as indicated. The example above now goes like this:

- In the first round, your opponent will choose Rock (A), and you need the round to end in a draw (Y), so you also choose Rock. This gives you a score of 1 + 3 = 4.
- In the second round, your opponent will choose Paper (B), and you choose Rock so you lose (X) with a score of 1 + 0 = 1.
- In the third round, you will defeat your opponent's Scissors with Rock for a score of 1 + 6 = 7.

Now that you're correctly decrypting the ultra top secret strategy guide, you would get a total score of 12.

Following the Elf's instructions for the second column, what would your total score be if everything goes exactly according to your strategy guide?
```

Remember:
* Each line from the input represents a round
* The format of each line is an A, B, or C in the first column, a space in the second column, and an X, Y, or Z in the third column
* A, B, C represent our opponent playing Rock, Paper, Scissors respectively
* X, Y, Z represent a desired outcome of loss, draw, and win respectively
* We need to determine whether we would play Rock, Paper, or Scissors to get the desired outcome
* You score 1 point for playing Rock, 2 for Paper, and 3 for Scissors
* In addition, you score 6 points for winning a round, 3 points for a draw, and 0 for a loss

## OpenGPT response

Go:

```go
package main

import (
  "bufio"
  "fmt"
  "os"
)

type Round struct {
  opponentMove string
  expectedEnd  string
  score        int
}

func (r *Round) move() string {
  switch r.opponentMove {
  case "A":
    switch r.expectedEnd {
    case "X":
      return "C"
    case "Y":
      return "A"
    case "Z":
      return "B"
    }
  case "B":
    switch r.expectedEnd {
    case "X":
      return "A"
    case "Y":
      return "B"
    case "Z":
      return "C"
    }
  case "C":
    switch r.expectedEnd {
    case "X":
      return "B"
    case "Y":
      return "C"
    case "Z":
      return "A"
    }
  }

  return ""
}

func (r *Round) updateScore() {
  switch r.move() {
  case "A":
    r.score = 1
  case "B":
    r.score = 2
  case "C":
    r.score = 3
  }

  switch r.expectedEnd {
  case "X":
    r.score += 0
  case "Y":
    r.score += 3
  case "Z":
    r.score += 6
  }
}

func main() {
  rounds := make([]Round, 0)
  score := 0

  f, _ := os.Open("input.txt")
  defer f.Close()

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    line := scanner.Text()
    opponentMove := string(line[0])
    expectedEnd := string(line[2])
    round := Round{opponentMove, expectedEnd, 0}
    round.updateScore()
    rounds = append(rounds, round)
    score += round.score
  }

  fmt.Println(score)
}
```

Rust:

```rust
use std::fs::File;
use std::io::{BufRead, BufReader};

struct Round {
    opponent_move: String,
    expected_end: String,
    score: i32,
}

impl Round {
    fn move(&self) -> String {
        match self.opponent_move.as_str() {
            "A" => match self.expected_end.as_str() {
                "X" => "C".to_string(),
                "Y" => "A".to_string(),
                "Z" => "B".to_string(),
                _ => "".to_string(),
            },
            "B" => match self.expected_end.as_str() {
                "X" => "A".to_string(),
                "Y" => "B".to_string(),
                "Z" => "C".to_string(),
                _ => "".to_string(),
            },
            "C" => match self.expected_end.as_str() {
                "X" => "B
```
