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
