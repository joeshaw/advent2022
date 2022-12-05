package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  f, _ := os.Open("input.txt")
  defer f.Close()

  scanner := bufio.NewScanner(f)

  sum := 0
  for scanner.Scan() {
    line := scanner.Text()
    firstCompartment := line[:len(line)/2]
    secondCompartment := line[len(line)/2:]
    for _, char := range firstCompartment {
      if strings.ContainsRune(secondCompartment, char) {
        if char >= 'a' && char <= 'z' {
          sum += int(char - 'a' + 1)
        } else {
          sum += int(char - 'A' + 27)
        }
        break
      }
    }
  }

  fmt.Println(sum)
}
