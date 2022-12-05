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
  group := make([]string, 0)
  for scanner.Scan() {
    line := scanner.Text()
    group = append(group, line)
    if len(group) == 3 {
      for _, char := range group[0] {
        if strings.ContainsRune(group[1], char) && strings.ContainsRune(group[2], char) {
          if char >= 'a' && char <= 'z' {
            sum += int(char - 'a' + 1)
          } else {
            sum += int(char - 'A' + 27)
          }
          break
        }
      }
      group = make([]string, 0)
    }
  }

  fmt.Println(sum)
}
