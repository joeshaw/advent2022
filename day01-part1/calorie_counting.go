package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
