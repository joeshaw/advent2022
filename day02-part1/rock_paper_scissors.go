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
