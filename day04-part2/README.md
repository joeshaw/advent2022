# Day 4, part 2

ChatGPT didn't get this right on the first try: it was counting the overlapped sections rather than just the number of overlapping pairs, but simply regenerating it gave a correct program the second time.

Run the generated code with `python3 cleanup.py`.

## Prompt

Moving on to part 2, which builds upon the earlier part.  As before, generate a complete program to solve the puzzle.

Part 2 prompt:
```
It seems like there is still quite a bit of duplicate work planned. Instead, the Elves would like to know the number of pairs that overlap at all.

In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't overlap, while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6, and 2-6,4-8) do overlap:

- 5-7,7-9 overlaps in a single section, 7.
- 2-8,3-7 overlaps all of the sections 3 through 7.
- 6-6,4-6 overlaps in a single section, 6.
- 2-6,4-8 overlaps in sections 4, 5, and 6.

So, in this example, the number of overlapping assignment pairs is 4.

In how many assignment pairs do the ranges overlap?
```

## ChatGPT response

Here is a solution in Python that builds upon the earlier solution to solve the second part of the puzzle:

```python
def overlaps(a, b):
    a_start, a_end = map(int, a.split("-"))
    b_start, b_end = map(int, b.split("-"))
    return a_start <= b_end and b_start <= a_end

count = 0
with open("input.txt") as f:
    for line in f:
        a, b = line.strip().split(",")
        if overlaps(a, b):
            count += 1
print(count)
```
