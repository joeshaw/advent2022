# Day 6, part 2

This carries on the ChatGPT conversation from part 1.

Run the generated code with `python3 tuner.py`.

## Prompt

Good job!  Now we'll move onto part 2 of the puzzle, which builds upon part 1.  As before, generate a complete program that solves the puzzle.

Here is the part two prompt:
```
Your device's communication system is correctly detecting packets, but still isn't working. It looks like it also needs to look for messages.

A start-of-message marker is just like a start-of-packet marker, except it consists of 14 distinct characters rather than 4.

Here are the first positions of start-of-message markers for all of the above examples:

- mjqjpqmgbljsphdztnvjfqwrcgsmlb: first marker after character 19
- bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 23
- nppdvjthqldpwncqszvftbrmjlhg: first marker after character 23
- nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 29
- zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 26

How many characters need to be processed before the first start-of-message marker is detected?
```

## ChatGPT response

Here is a solution in Python that solves part 2 of the puzzle:

```python
def main():
    with open('input.txt') as f:
        data = f.read().strip()
        chars = []
        for i, c in enumerate(data):
            chars.append(c)
            if len(chars) > 14:
                chars.pop(0)
            if len(chars) == 14 and len(set(chars)) == 14:
                print(i+1)
                break

if __name__ == '__main__':
    main()
```

This solution is similar to the solution for part 1, but it maintains a list of the 14 most recent characters seen so far, instead of 4. It also looks for a list of 14 distinct characters rather than 4, and prints the current position plus 1 when it finds such a list.
