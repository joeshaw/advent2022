# Day 6, part 1

I _loved_ the way ChatGPT solved this one.  A moving window of characters, and using a `set` to detect if there were duplicates.  Honestly this is probably more elegant than I would have written on my first try.

Run the generated code with `python3 tuner.py`.

## Prompt

Advent of Code is a Christmas-themed series of programming puzzles.  Each day there are two parts to a puzzle, with the second part building upon the first one.  I will provide you with the prompt for the first part.  I want you to generate a complete program in Go, Rust, or Python to solve the puzzle.  The program should read its input from a file named `input.txt`.

Here's the prompt:
```
The preparations are finally complete; you and the Elves leave camp on foot and begin to make your way toward the star fruit grove.

As you move through the dense undergrowth, one of the Elves gives you a handheld device. He says that it has many fancy features, but the most important one to set up right now is the communication system.

However, because he's heard you have significant experience dealing with signal-based systems, he convinced the other Elves that it would be okay to give you their one malfunctioning device - surely you'll have no problem fixing it.

As if inspired by comedic timing, the device emits a few colorful sparks.

To be able to communicate with the Elves, the device needs to lock on to their signal. The signal is a series of seemingly-random characters that the device receives one at a time.

To fix the communication system, you need to add a subroutine to the device that detects a start-of-packet marker in the datastream. In the protocol being used by the Elves, the start of a packet is indicated by a sequence of four characters that are all different.

The device will send your subroutine a datastream buffer (your puzzle input); your subroutine needs to identify the first position where the four most recently received characters were all different. Specifically, it needs to report the number of characters from the beginning of the buffer to the end of the first such four-character marker.

For example, suppose you receive the following datastream buffer:

mjqjpqmgbljsphdztnvjfqwrcgsmlb

After the first three characters (mjq) have been received, there haven't been enough characters received yet to find the marker. The first time a marker could occur is after the fourth character is received, making the most recent four characters mjqj. Because j is repeated, this isn't a marker.

The first time a marker appears is after the seventh character arrives. Once it does, the last four characters received are jpqm, which are all different. In this case, your subroutine should report the value 7, because the first start-of-packet marker is complete after 7 characters have been processed.

Here are a few more examples:

- bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 5
- nppdvjthqldpwncqszvftbrmjlhg: first marker after character 6
- nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 10
- zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 11

How many characters need to be processed before the first start-of-packet marker is detected?
```

## ChatGPT response

Here is a solution in Python:

```python
def main():
    with open('input.txt') as f:
        data = f.read().strip()
        chars = []
        for i, c in enumerate(data):
            chars.append(c)
            if len(chars) > 4:
                chars.pop(0)
            if len(chars) == 4 and len(set(chars)) == 4:
                print(i+1)
                break

if __name__ == '__main__':
    main()
```

This solution works by reading the input data from the `input.txt` file, and then iterating over the characters in the data. For each character, it maintains a list of the four most recent characters seen so far. If the list ever contains four characters that are all different, it prints the current position in the data plus 1 (since the position is 0-indexed) and then exits.
