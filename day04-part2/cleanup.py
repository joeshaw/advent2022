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
