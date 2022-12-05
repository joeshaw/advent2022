# Day 1, part 2

I continued the OpenGPT conversation from part 1, which allows it to use context from the previous task.

Run the generated code with `python3 calorie_counting.py < input.txt`.

## Prompt

We will now move onto part 2 of today's puzzle.  As before, generate a program that will solve this puzzle given the prompt below.  You may reuse and modify your code from the previous answer to solve it.

The prompt:
```
By the time you calculate the answer to the Elves' question, they've already realized that the Elf carrying the most Calories of food might eventually run out of snacks.

To avoid this unacceptable situation, the Elves would instead like to know the total Calories carried by the top three Elves carrying the most Calories. That way, even if one of those Elves runs out of snacks, they still have two backups.

In the example above, the top three Elves are the fourth Elf (with 24000 Calories), then the third Elf (with 11000 Calories), then the fifth Elf (with 10000 Calories). The sum of the Calories carried by these three elves is 45000.

Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
```

## OpenGPT response

Here is a possible solution to the problem in Python:

```python
import sys
from collections import Counter

# Read the input from standard input
lines = [line.strip() for line in sys.stdin]

# Initialize a counter to count the Calories of each Elf
calorie_counts = Counter()

# Keep track of the current Elf and the total number of Calories for that Elf
current_elf = None
total_calories = 0

for line in lines:
    if line == "":
        # If we encounter a blank line, add the current Elf and their Calories
        # to the counter, and reset the total number of Calories
        calorie_counts[current_elf] = total_calories
        current_elf = None
        total_calories = 0
    else:
        # If we encounter a number, it represents the Calories of a food item
        # for the current Elf
        calories = int(line)
        total_calories += calories

        # If the current Elf is not set, this is the first food item for the Elf
        # so we set the current Elf to the number of Calories in the food item
        if current_elf is None:
            current_elf = calories

# Add the final Elf and their Calories to the counter
calorie_counts[current_elf] = total_calories

# Get the top three Elves with the most Calories
top_elves = calorie_counts.most_common(3)

# Calculate the total number of Calories for the top three Elves
total_calories = sum(elf[1] for elf in top_elves)

# Print the total number of Calories
print(total_calories)
```

This program uses a `Counter` from the `collections` module to count the number of Calories carried by each Elf. It then uses the `most_common` method of the `Counter` to get the top three Elves with the most Calories. Finally, it calculates the total number of Calories carried by the top three Elves and prints the result.

You can run this program by saving the code to a file (e.g. `calorie_counting.py`), and running the following command:

```
python3 calorie_counting.py
```

You can then provide the input to the program through standard input, for example:

```
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
```

This should produce the output `45000`, indicating that the top three Elves are carrying a total of 45000 Calories.
