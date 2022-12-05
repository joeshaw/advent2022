# Day 3, part 1

Continuing the conversation from Day 2, part 2.  ChatGPT produced both Go and Rust versions that are complete and correct.

## Prompt

Moving onto the next day's puzzles.  Part 1 is given in the prompt below.  As before, you must generate a complete program and may use any of the pre-approved programming languages.  The input is given in an `input.txt` file.

The prompt:
```
One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.

Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.

The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer to different types of items).

The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

For example, suppose you have the following list of contents from six rucksacks:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

- The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
- The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
- The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
- The fourth rucksack's compartments only share item type v.
- The fifth rucksack's compartments only share item type t.
- The sixth rucksack's compartments only share item type s.

To help prioritize item rearrangement, every item type can be converted to a priority:

- Lowercase item types a through z have priorities 1 through 26.
- Uppercase item types A through Z have priorities 27 through 52.

In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?
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
```

Rust:

```rust
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let f = File::open("input.txt").expect("file not found");
    let reader = BufReader::new(f);

    let mut sum = 0;
    for line in reader.lines() {
        let line = line.expect("unable to read line");
        let first_compartment = &line[..line.len()/2];
        let second_compartment = &line[line.len()/2..];

        for ch in first_compartment.chars() {
            if second_compartment.contains(ch) {
                if ch >= 'a' && ch <= 'z' {
                    sum += (ch as u8 - b'a' + 1) as i32;
                } else {
                    sum += (ch as u8 - b'A' + 27) as i32;
                }
                break;
            }
        }
    }

    println!("{}", sum);
}
```
