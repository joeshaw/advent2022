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
