# Day 3, part 2

Continuing the conversation from Day 3, part 1.  ChatGPT produced a correct Go program and an incomplete Rust program.

Run the generated code with `go run rucksack.go`.

## Prompt

Continuing to part 2 of today's puzzle.  As before, you must generate a complete program and may use any of the pre-approved programming languages.  The input is given in an `input.txt` file.

The prompt:
```
As you finish identifying the misplaced items, the Elves come to you with another issue.

For safety, the Elves are divided into groups of three. Every Elf carries a badge that identifies their group. For efficiency, within each group of three Elves, the badge is the only item type carried by all three Elves. That is, if a group's badge is item type B, then all three Elves will have item type B somewhere in their rucksack, and at most two of the Elves will be carrying any other item type.

The problem is that someone forgot to put this year's updated authenticity sticker on the badges. All of the badges need to be pulled out of the rucksacks so the new authenticity stickers can be attached.

Additionally, nobody wrote down which item type corresponds to each group's badges. The only way to tell which item type is the right one is by finding the one item type that is common between all three Elves in each group.

Every set of three lines in your list corresponds to a single group, but each group can have a different badge item type. So, in the above example, the first group's rucksacks are the first three lines:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg

And the second group's rucksacks are the next three lines:

wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

In the first group, the only item type that appears in all three rucksacks is lowercase r; this must be their badges. In the second group, their badge item type must be Z.

Priorities for these items must still be found to organize the sticker attachment efforts: here, they are 18 (r) for the first group and 52 (Z) for the second group. The sum of these is 70.

Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the priorities of those item types?
```

## ChatGPT response

Go:

```go
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
```

Rust:

```rust
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let f = File::open("input.txt").expect("file not found");
    let reader = BufReader::new(f);

    let mut sum = 0;
    let mut group = Vec::new();
    for line in reader.lines() {
        let line = line.expect("unable to read line");
        group.push(line);
        if group.len() == 3 {
            for ch in group[0].chars() {
                if group[1].contains(ch) && group[2].contains(ch) {
                    if ch >= 'a' && ch <= 'z' {
                        sum += (ch as u8 - b'a' + 1) as
```
